test:
	go test -cover -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o coverage.html

test-debug:
	go test -cover -v -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o coverage.html
