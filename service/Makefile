SRCDIR=./cmd
GATEWAY=gateway
SERVICE=service

.PHONY: all
all: gateway service

.PHONY: gateway
gateway: tidy
	go build -o $@ $(SRCDIR)/$(GATEWAY)/main.go

.PHONY: service
service: tidy
	go build -o $@ $(SRCDIR)/$(SERVICE)/main.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: proto
proto:
	protoc -I ./proto -I ./proto/google -I ./proto/openapiv2 \
		--go_out ./gen --go_opt paths=source_relative \
		--go-grpc_out ./gen --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./gen --grpc-gateway_opt paths=source_relative \
		--openapiv2_out ./gen \
		./proto/service.proto

.PHONY: clean
clean:
	$(RM) $(SERVICE) $(GATEWAY)
