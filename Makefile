dist ?= ./build
folder ?= ./internal
count ?= 1
time ?= "1s"
test ?= "."

all: install run

install:
	go install github.com/go-task/task/v3/cmd/task@latest && \
	go install golang.org/x/tools/gopls@latest && \
	go install github.com/marco-souza/hooker@latest && hooker init

run: main.go
	go run ${flags} main.go

build: main.go
	go build -o ${dist}/pkg ${flags} main.go

fmt:
	go fmt ./...

test:
	go test ${flags} ./...

bench: tests/bench/...
	go test -bench=${test} ./tests/bench/... -count=${count} -benchmem -benchtime=${time}

encrypt: .env
	@go run main.go encrypt .env
	@sed -e 's/=.*/=""/g' .env > .env.example # TODO: implement in pkg

decrypt: .env.gpg
	@go run main.go decrypt .env

# TODO: implement in pkg
gen:
	@go run ./cmd/cli/cli.go create ${folder} ${name}
