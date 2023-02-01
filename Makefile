all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/SignatureLetter.git

pull:
	git pull git@github.com:RB-PRO/SignatureLetter.git

pushW:
	git pushW https://github.com/RB-PRO/SignatureLetter.git

pullW:
	git pull https://github.com/RB-PRO/SignatureLetter.git

pushCar:
	scp main root@194.87.107.129:go/SignatureLetter/

build-config:
	go env GOOS GOARCH

build-linux-osx:
	export GOARCH=arm
	export GOOS=linux
	go env GOOS GOARCH
	go build ./cmd/main/main.go  

build-linux-linux:
	export GOARCH=amd64
	export GOOS=linux
	go env GOOS GOARCH
	go build ./cmd/main/main.go

build-linux-windows:
	export GOARCH=amd64
	export GOOS=windows
	go env GOOS GOARCH
	go build ./cmd/main/main.go