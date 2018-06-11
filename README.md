# drunkmbot

One and only (yeah, right) Discord bot written in Go. Project started because of reasons and will to learn
new language.

*WARNING*

This might work only for Linux. Mac/Windows ain't supported

## Build

### Executable

Just build `cmd/drunkmbot.go`

`go build cmd/drunkmbot.go`

### Plugins

Go to each plugin dir and build in pluginmode

`go build -buildmode=plugin *.go`
