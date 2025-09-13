
build:
	CGO_ENABLED=1 go build .

test:
	CGO_ENABLED=1 go test -v .
bench:
	CGO_ENABLED=1 go test -v -benchmem -bench .

lint:
	gofmt -w -s *.go


