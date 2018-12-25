ENTRYPOINT=./cmd/nvu/cli.go
OUTPUT=nvu

build:
	go build -o $(OUTPUT) $(ENTRYPOINT)

debug-release: build
	mv $(OUTPUT) ~/local/bin

run:
	go run $(ENTRYPOINT) package.json
