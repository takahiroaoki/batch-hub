# build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o batch-hub

# test
ut:mockgen
	go clean -testcache \
	&& go test -v ./...

mockgen:
	rm -f ./**/*_mock.go
	mockgen -source=./config/database.go -destination=./config/database_mock.go -package=config

# others
lint: mockgen
	golangci-lint run