SRCDIR=./cmd
TARGET=configurator

.PHONY: build
build: proto tidy
	go build -o $(TARGET) $(SRCDIR)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: proto
proto:
	protoc -I ./proto -I ./proto/google -I ./proto/openapiv2 \
		--go_out ./proto --go_opt paths=source_relative \
		--go-grpc_out ./proto --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
		--openapiv2_out . \
		./proto/service.proto

.PHONY: clean
clean:
	$(RM) $(TARGET)
