# Networking Tools
Golang / Python Networking / Pentesting Tools


|Name|Link|About|
| :------------- | :----------: | -----------: |
|  cli-gobuster |   [Repo](https://github.com/self-rep/networktools/tree/main/cli-gobuster) | A Worse Version Of Gobuster Pentesting Tool |


# CLI-GOBUSTER
## Install
- git clone https://github.com/self-rep/networktools
- cd networktools
- go build -o gobuster.exe -ldflags="-s -w" main.go
## Usage
- .\gobuster.exe -h (Shows Help Menu)
- .\gobuster.exe -host=http://example.com -port=<PORT> -wordlist=<PATH TO WORDLIST> -output=true -timeout=<TIMEOUT IN MS/SECONDS>
