package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// 1. Get Gemini API Key from System Environment Variables
	// You need to set this in your terminal before running the program
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("[!] Error: GEMINI_API_KEY environment variable is not set.")
		fmt.Println("Please get an API key from Google AI Studio and set it using:")
		fmt.Println("export GEMINI_API_KEY='your_key_here'")
		return
	}

	// 2. Initialize the Gemini AI Client
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("[!] Failed to create Gemini client: %v", err)
	}
	defer client.Close()

	// Using the recommended fast and smart model
	model := client.GenerativeModel("gemini-1.5-flash")

	// 3. Get input from the user about what interface they want to build
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("==================================================")
	fmt.Println("        AI Real-Time Linux GUI Builder            ")
	fmt.Println("==================================================")
	fmt.Print("Describe the interface you want to build:\n> ")
	
	userPrompt, _ := reader.ReadString('\n')
	userPrompt = strings.TrimSpace(userPrompt)

	if userPrompt == "" {
		fmt.Println("[!] Input cannot be empty.")
		return
	}

	// 4. System Prompt engineering to force Gemini to return ONLY pure executable Python code
	systemInstruction := `You are an expert Python developer specialized in CustomTkinter. 
Your task is to generate complete, working, and professional Python GUI code based on the user's request.
CRITICAL RULES:
1. Return ONLY the raw Python code. Do NOT include markdown code blocks like ` + "```python" + ` or ` + "
```" + `.
2. Do NOT include any explanations, text, or comments outside the code.
3. Make the UI modern, well-spaced, professional, and functional.
4. Ensure 'import customtkinter as ctk' is used and 'app.mainloop()' is called at the end.`

	fmt.Println("\n[*] Consulting AI Architect to generate code...")

	// 5. Send prompt to Gemini
	resp, err := model.GenerateContent(ctx, genai.Text(systemInstruction), genai.Text(userPrompt))
	if err != nil {
		log.Fatalf("[!] Error generating content from Gemini: %v", err)
	}

	// 6. Extract the generated code from response
	var generatedCode string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				generatedCode += fmt.Sprintf("%v", part)
			}
		}
	}

	// Clean up formatting just in case the AI added backticks
	generatedCode = strings.TrimPrefix(generatedCode, "```python")
	generatedCode = strings.TrimPrefix(generatedCode, "
```")
	generatedCode = strings.TrimSuffix(generatedCode, "```")
	generatedCode = strings.TrimSpace(generatedCode)

	if generatedCode == "" {
		fmt.Println("[!] Failed to generate valid code from the prompt.")
		return
	}

	// 7. Write the dynamic code into a local python file
	outputFile := "dynamic_interface.py"
	err = os.WriteFile(outputFile, []byte(generatedCode), 0644)
	if err != nil {
		log.Fatalf("[!] Failed to save code to file: %v", err)
	}
	fmt.Printf("[+] AI compiled code successfully saved to '%s'\n", outputFile)

	// 8. Launch the generated interface live on Linux
	fmt.Println("[*] Launching your custom interface now...")
	cmd := exec.Command("python3", outputFile)
	
	// Running it in background so it doesn't freeze the terminal
	err = cmd.Start()
	if err != nil {
		log.Fatalf("[!] Failed to execute the generated interface: %v", err)
	}
	fmt.Println("[✓] Done! Check your desktop screen.")
}
