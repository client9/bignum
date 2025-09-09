
build:
	CGO_ENABLED=1 go build ./...

test:
	go test -v .
bench:
	CGO_ENABLED=1 go test -v -benchmem -bench .

