# test
ut:mockgen
	go clean -testcache \
	&& go test -v ./...

mockgen:
	rm -f ./**/*_mock.go
	mockgen -source=./config/database.go -destination=./config/database_mock.go -package=config