build:
	go build -o /usr/local/bin/go-work .

run: build
	go-work