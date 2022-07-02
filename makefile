config = config.yaml
Path  = ./web
Exec  = ./web/main.go
run:${Exec}
	go run ${Exec} ${config}

clean:
	rm -rf build