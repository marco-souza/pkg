dist ?= ./build
folder ?= ./internal
count ?= 1
time ?= "1s"
test ?= "."

all: install run

i: install
install:
	go install golang.org/x/tools/gopls@latest && \
	go install github.com/marco-souza/hooker@latest && hooker init && \
	go install .

run: main.go
	@go run main.go

fmt:
	go fmt ./...

t: test
test:
	go test ./...

encrypt: .env
	@go run main.go encrypt .env
	@sed -e 's/=.*/=""/g' .env > .env.example # TODO: implement in pkg

decrypt: .env.gpg
	@go run main.go decrypt .env
