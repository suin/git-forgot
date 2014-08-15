run: fmt
	go run *.go --debug ${args}

fmt:
	go fmt ./...

build: fmt
	go build -o git-forgot *.go
