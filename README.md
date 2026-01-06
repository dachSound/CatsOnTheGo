A simple GO program that downloads cat pictures from cataas.com (Cats as a service)

What it does:
	
	Downloads cat pictures every 1.5 Seconds (Can be changed with the delay variable)
	Saves them to cat/ directory
	Names files sequentially: cat0.jpg, cat1.jpg, etc

Prerequisites: GO Installation

Quick Install (Recommended)

Choose your operating system:

Mac

			# Using Homebrew (recommended)
			brew install go

			# Or download from official site:
			# https://go.dev/dl/

Windows

			# Using Winget (Windows 11)
			winget install GoLang.Go

			# Or download installer from:
			# https://go.dev/dl/

Linux (Ubuntu/Debian)

			# Using apt
			sudo apt update
			sudo apt install golang-go

			# Or for latest version:
			sudo snap install go --classic

Linux (Fedora/RHEL)

			sudo dnf install golang

Verify Installation

			go version




How to install and run CatsOnTheGo: 

	```bash git clone https://github.com/dachSound/CatsOnTheGo
			cd CatsOnTheGo
			go run main.go
	```

Optional: Build an Executable

		# Build for your current OS
		go build -o catsonthego .

		# Run the executable
		./catsonthego
