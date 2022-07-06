config = config.yaml
Path  = ./web
Exec  = ./web/main.go

init:
	go mod tidy

run:${Exec}
	go run ${Exec} ${config}

clean:
	rm -rf build