build:
	go build -o bin/$(basename $(pwd))

run:
	./bin/app

test:
	go test -v ./... -count=1 # --race