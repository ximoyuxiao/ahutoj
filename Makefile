config = config.yaml
Path  = ./web
Exec  = ./web/main.go
tagert = ./tmp/bin/main
all:init run

init:
	go mod tidy

run:${Exec}
	air -c air.conf

build:${Exec}
	go build -o ${tagert} ${Exec}

clean:
	rm -rf ${tagert}