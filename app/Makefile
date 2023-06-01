build:
	go build -tags netgo -o bin/$(basename $(pwd))
	# CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

run:
	./bin/app

test:
	go test -v ./... -count=1 # --race