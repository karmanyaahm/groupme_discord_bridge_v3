flags = CGO_ENABLED=0 GOOS=linux
targ = builds/groupme_discord_bridge_
ldflags = -ldflags="-s" 
#ldflag -w ?
upx_command = echo 
#MAKEFLAGS += -Otarget

all: windows_amd64 linux_amd64 linux_arm64 linux_armv6 linux_armv7 linux_i386

release:
	$(MAKE) clean
	$(MAKE) upx all 
	$(MAKE) checksum
	
amd64:
	$(flags) GOARCH=amd64 go build -o $(targ)amd64 ./main.go
upx:
	$(eval upx_command=upx $(targ))
checksum:
	cd builds; \
	  sha256sum * > sha256

linux_amd64:
	$(flags) GOARCH=amd64 GOOS=linux go build $(ldflags) -o $(targ)$@ ./main.go
	$(upx_command)$@
linux_arm64:
	$(flags) GOARCH=arm64 GOOS=linux go build $(ldflags) -o $(targ)$@      ./main.go
	$(upx_command)$@

linux_i386:
	$(flags) GOARCH=386 GOOS=linux go build $(ldflags) -o $(targ)$@      ./main.go
	$(upx_command)$@
linux_armv6:
	$(flags) GOARCH=arm GOARM=6 GOOS=linux go build $(ldflags) -o $(targ)$@ ./main.go
	$(upx_command)$@
linux_armv7:
	$(flags) GOARCH=arm GOARM=7 GOOS=linux go build $(ldflags) -o $(targ)$@ ./main.go
	$(upx_command)$@
windows_amd64:
	$(flags) GOARCH=amd64 GOOS=windows go build $(ldflags) -o $(targ)$@ ./main.go
	$(upx_command)$@
clean:
	rm -rf builds/
	go clean


