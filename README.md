# 🚀 Prompt2UI-Linux

An advanced, real-time AI GUI Compiler written in **Go (Golang)**. It acts as an autonomous AI agent that accepts natural language instructions from the Linux terminal, processes them via the **Google Gemini API**, dynamically generates production-ready **CustomTkinter (Python)** code on the fly, and instantly executes it as a live native application.

---

## ✨ Features

* **Pure Dynamic Generation:** Unlike rigid rule-based tools, this system writes custom code from scratch based on any complex text prompt.
* **Modern Dark-Theme UI:** Automatically structures code using `CustomTkinter` for sleek, responsive, and professional user interfaces.
* **Seamless Go & Python Integration:** Combines the blazing-fast execution and system-level control of Go with the rich GUI capabilities of Python.
* **Isolated Script Execution:** Automatically compiles, saves, and launches the generated frontend cleanly in the Linux background.

---

## 🏗️ Architecture Flow

1. **User Input:** Enter a prompt in the Go TUI (e.g., *"Create a red dark-themed network monitor with 3 status gauges"*).
2. **AI Synthesis:** Go formats the prompt with strict system engineering constraints and transmits it to `gemini-1.5-flash`.
3. **Code Extraction:** The AI streams back sanitized, raw Python code with no markdown wrapper clutter.
4. **Live Deployment:** The Go runtime writes the payload into `dynamic_interface.py` and immediately invokes `python3` to spawn the interface.

---

## 🛠️ Prerequisites & Installation

Ensure you have **Go** and **Python 3** installed on your Linux machine.

### 1. Install Dependencies
```bash
# Update system packages
sudo apt update

# Install Golang and Python3 pip (Ubuntu/Debian)
sudo apt install golang python3-pip -y

# Install CustomTkinter UI framework
pip install customtkinter
```
### 2. Get the Go API Libraries
```Bash
go get [github.com/google/generative-ai-go/genai](https://github.com/google/generative-ai-go/genai)
go get google.golang.org/api/option
```
## 🚀 Setup & Execution
### 1. Clone the Repository
```
git clone [https://github.com/rushyaayt/Prompt2UI-Linux.git](https://github.com/rushyaayt/Prompt2UI-Linux.git)
cd Prompt2UI-Linux
```
### 2. Set Up Your Gemini API Key
Obtain a free API key from Google AI Studio and expose it to your environment:
```Bash
export GEMINI_API_KEY="your_actual_api_key_here"
```
### 3. Run the Compiler
```
go run pro2ui.go
```
## 💡 Prompt Examples to Try
Once the prompt screen loads, try feeding it these descriptions to witness real-time compilation:
- 💥 Create a dark-theme cyber security dashboard with an 'Initiate Scan' button and a terminal log console.
- 💥 Build a clean system performance monitor window showing placeholder bars for CPU, RAM, and Network usage.
- 💥 Generate an enterprise-grade login portal with email/password input fields and a blue authentication button.


## 📁 Project Structure
``` Plaintext
Prompt2UI-Linux/
├── main.go               # Core compiler engine (Go)
├── go.mod                # Dependency tracking file
├── .gitignore            # Keeps the repo clean from compiled cache
└── README.md             # Documentation (This file)
```
## 🤝 Contributing
Contributions, issues, and feature requests are welcome! Feel free to check the issues page if you want to expand this into supporting web or multi-window structures.
## 📜 License
Distributed under the MIT License. See LICENSE for more information.
