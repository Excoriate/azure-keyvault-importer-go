fmt:
	gofmt -s -w .
test:
	go test ./... -cover
lint:
	golangci-lint run
build:
	go build -v -o importer_keyvault main.go
run:
	chmod +x importer_keyvault; ./importer_keyvault
