install:
	go install github.com/air-verse/air@latest

run:
	-mkdir -p ./bin
	air --build.cmd "go build -o ./bin/server ./cmd/server" --build.bin "./bin/server"
