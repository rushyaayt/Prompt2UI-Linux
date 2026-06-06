# Prompt2UI-Linux
## 🛠️ First install these two dependencies
```
pip install customtkinter
```
```
sudo apt install python3 python3-pip -y
```
```
go get github.com/google/generative-ai-go/genai
go get google.golang.org/api/option


```
## 💻 Go Source Code (main.go)
This code is completely dynamic. The interface will be written and run by AI as you tell it:
## 🛠️ How to set it up?
Since this is a real AI project, it will require a (Free) API Key.
- 1 Go to Google AI Studio and get your free Gemini API Key with one click.
- 2 Set that key in your Linux terminal like this:
export GEMINI_API_KEY="Put your API key here"
```
go run pro2ui.go
```
## 🗣️ You can call it anything now!
When the terminal displays Describe the interface you want to build:, try typing:
- Example 1:
- ```
  Create a red dark-themed hacking tool dashboard with 4 danger buttons and a big status text block.
  ```
- Example 2:
- ```
  Create a simple calculator with a blue theme and circular buttons.
  ```
- Example 3:
- ```
  Build a system monitor screen showing dummy CPU and RAM bars.
  ```
This tool will fetch new code from the internet every time it is written, create its dynamic_interface.py file, and open it on the screen before your eyes!
### Real-Time AI GUI Compiler written in Go. It accepts natural language instructions, acts as an AI agent using Gemini API to write production-ready CustomTkinter Python code on-the-fly, and instantly executes it as a standalone Linux application.
