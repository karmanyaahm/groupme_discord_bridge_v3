flags = CGO_ENABLED=0 GOOS=linux
targ = builds/groupme_discord_bridge_linux_
ldflags = -ldflags="-s" 
#ldflag -w ?

amd64:
	$(flags) GOARCH=amd64 go build -o $(targ)x86_64 ./main.go

all:
	$(flags) GOARCH=amd64 go build $(ldflags) -o $(targ)x86_64 ./main.go
	$(flags) GOARCH=arm64 go build $(ldflags) -o $(targ)arm64      ./main.go
	$(flags) GOARCH=386 go build $(ldflags) -o $(targ)i386      ./main.go
	$(flags) GOARCH=arm GOARM=6 go build $(ldflags) -o $(targ)armv6 ./main.go
	$(flags) GOARCH=arm GOARM=7 go build $(ldflags) -o $(targ)armv7 ./main.go
	upx $(targ)*



