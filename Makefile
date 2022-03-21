test:
	go test -cover -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o coverage.html

test-debug:
	go test -cover -v -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o coverage.html

proto:
	@echo "Generating Proto Data"
	@for filename in *.proto; do \
		mkdir -p pb/$${filename%.*}$f ;\
		protoc --proto_path=. --go_out=pb/$${filename%.*} --go_opt=paths=source_relative --go-grpc_out=pb/$${filename%.*} --go-grpc_opt=paths=source_relative $$filename ; \
	done

clean:
	rm -R pb/*
