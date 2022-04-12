PROJECT:=re-chinese

.PHONY: build
build:
	CGO_ENABLED=0 go build -o re-chinese main.go
build-linux:
	env GOOS=linux GOARCH=amd64 go build
#?.PHONY: test
#?test:
#?	go test -v ./... -cover

#?.PHONY: docker
#?docker:
#?	docker build . -t re-chinese:latest
