setup go

important setups: https://sourabhbajaj.com/mac-setup/Homebrew/Cask.html

- install go
- set env variable by updating .bash_profile (open this file to see added)
	`export GOPATH=$HOME/workspace/go-workspace` # Default go workspace
	`export GOROOT=/usr/local/opt/go/libexec`
	`export PATH=$PATH:$GOPATH/bin`
	`export PATH=$PATH:$GOROOT/bin`

	$GOPATH/src : Where your Go projects / programs are located
	$GOPATH/pkg : contains every package objects
	$GOPATH/bin : The compiled binaries home

- reload bash_profile with: command `. .bash_profile` or `source ~/.bash_profile`

- Import a Go package e.g gorilla/mux
` go get -u github.com/gorilla/mux` and installs in github.com folder

- symbolic link
`sudo ln -s /Applications/Visual\ Studio\ Code.app/Contents/Resources/app/bin/code /usr/local/bin/vscode`

//could not launch process: debugserver or lldb-server not found: install XCode's command line tools or lldb-server
`xcode-select  --install`
