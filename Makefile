test:
	@FILES="$(shell go list ./... | grep -v /pb/ | grep -v github.com/Tackem-org/Global/system/run)"; go test -cover -coverprofile=cover.out $$FILES
	@go tool cover -html=cover.out -o coverage.html

test-debug:
	@FILES="$(shell go list ./... | grep -v /pb/ | grep -v github.com/Tackem-org/Global/system/run)"; go test -cover -v -coverprofile=cover.out $$FILES
	@go tool cover -html=cover.out -o coverage.html

proto:
	@echo "Generating Proto Data..."
	@mkdir -p pb/config pb/regclient pb/registration pb/remoteweb pb/user pb/web
	@echo "Generating Config Proto Data..."
	@protoc --proto_path=protos --go_out=pb/config --go_opt=paths=source_relative --go-grpc_out=pb/config --go-grpc_opt=paths=source_relative protos/config.proto
	@echo "Generating RegClient Proto Data..."
	@protoc --proto_path=protos --go_out=pb/regclient --go_opt=paths=source_relative --go-grpc_out=pb/regclient --go-grpc_opt=paths=source_relative protos/regclient.proto
	@echo "Generating Registration Data..."
	@protoc --proto_path=protos --go_out=pb/registration --go_opt=paths=source_relative --go-grpc_out=pb/registration --go-grpc_opt=paths=source_relative protos/registration.proto
	@echo "Generating RemoteWeb Data..."
	@protoc --proto_path=protos --go_out=pb/remoteweb --go_opt=paths=source_relative --go-grpc_out=pb/remoteweb --go-grpc_opt=paths=source_relative protos/remoteweb.proto
	@echo "Generating User Data..."
	@protoc --proto_path=protos --go_out=pb/user --go_opt=paths=source_relative --go-grpc_out=pb/user --go-grpc_opt=paths=source_relative protos/user.proto
	@echo "Generating Web Data..."
	@protoc --proto_path=protos --go_out=pb/web --go_opt=paths=source_relative --go-grpc_out=pb/web --go-grpc_opt=paths=source_relative protos/web.proto
	@echo "Generated Proto Data"

clean:
	@rm -R pb/*
	@rm cover.out coverage.html
