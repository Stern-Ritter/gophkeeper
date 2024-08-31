.PHONY: build-server build-agent build

SERVER_VERSION := 0.8
SERVER_DIR := cmd/server
SERVER_OUTPUT := server

CLIENT_VERSION := 0.8
CLIENT_DIR := cmd/client
CLIENT_OUTPUT := client

BUILD_DATE = $(shell date +'%Y/%m/%d')
BUILD_COMMIT = $(shell git rev-parse HEAD)

CERTS_DIR=./certs
PRIVATE_KEY_PKCS8=private_pkcs8.pem
PRIVATE_KEY=private.pem
PUBLIC_KEY=public.pem
KEY_SIZE = 4096

DB_CONTAINER_NAME=db
POSTGRES_IMAGE=postgres:10
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres
POSTGRES_PORT=5432

STATIC_LINTER_NAME = multichecker
STATIC_LINTER_DIR = ./cmd/staticlint

gofmt:
	goimports -local github.com/Stern-Ritter/gophkeeper -w .

build-static-linter:
	@echo "Building $(STATIC_LINTER_NAME)..."
	go build -o $(STATIC_LINTER_DIR)/$(STATIC_LINTER_NAME) $(STATIC_LINTER_DIR)/$(STATIC_LINTER_NAME).go

lint: build-static-linter
	@echo "Running static analysis on the project..."
	$(STATIC_LINTER_DIR)/$(STATIC_LINTER_NAME) ./...

proto-gen:
	buf generate

openssl-tls-certs-gen:
	openssl req -x509 -newkey rsa:2048 -nodes -days 365 -keyout $(CERTS_DIR)/ca-key.pem -out $(CERTS_DIR)/ca-cert.pem -subj "/C=RU/ST=Russia/L=Moscow/O=DEV/OU=DEV/CN=CA/emailAddress=gophkeeper@yandex.ru"
	openssl req -new -keyout $(CERTS_DIR)/server-key.pem -out $(CERTS_DIR)/server-req.pem -config server-cert.cnf
	openssl x509 -req -in $(CERTS_DIR)/server-req.pem -CA $(CERTS_DIR)/ca-cert.pem -CAkey $(CERTS_DIR)/ca-key.pem -CAcreateserial -out $(CERTS_DIR)/server-cert.pem -days 365 -extfile server-cert.cnf -extensions req_ext
	openssl req -newkey rsa:2048 -nodes -keyout $(CERTS_DIR)/client-key.pem -out $(CERTS_DIR)/client-req.pem -subj "/C=RU/ST=Russia/L=Moscow/O=DEV/OU=DEV/CN=CA/emailAddress=gophkeeper@yandex.ru"
	openssl x509 -req -in $(CERTS_DIR)/client-req.pem -CA $(CERTS_DIR)/ca-cert.pem -CAkey $(CERTS_DIR)/ca-key.pem -CAcreateserial -out $(CERTS_DIR)/client-cert.pem -days 365

build-server:
	cd $(SERVER_DIR) && go build -buildvcs=false -ldflags "-X main.buildVersion=v$(SERVER_VERSION) -X main.buildDate=$(BUILD_DATE) -X main.buildCommit=$(BUILD_COMMIT)" -o $(SERVER_OUTPUT) && cd ../..

build-client:
	cd $(CLIENT_DIR) && \
	GOOS=windows GOARCH=amd64 go build -buildvcs=false -ldflags "-X main.buildVersion=v$(CLIENT_VERSION) -X main.buildDate=$(BUILD_DATE)" -o $(CLIENT_OUTPUT)_windows.exe && \
	GOOS=linux GOARCH=amd64 go build -buildvcs=false -ldflags "-X main.buildVersion=v$(CLIENT_VERSION) -X main.buildDate=$(BUILD_DATE)" -o $(CLIENT_OUTPUT)_linux && \
	GOOS=darwin GOARCH=amd64 go build -buildvcs=false -ldflags "-X main.buildVersion=v$(CLIENT_VERSION) -X main.buildDate=$(BUILD_DATE)" -o $(CLIENT_OUTPUT)_macos && \
	cd ../..

build: build-server build-client

clean:
	rm -f $(SERVER_DIR)/$(SERVER_OUTPUT)
	rm -f $(CLIENT_DIR)/$(CLIENT_OUTPUT)

run-db:
	docker run -d \
	    -e POSTGRES_USER=$(POSTGRES_USER) \
	    -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	    -e POSTGRES_DB=$(POSTGRES_DB) \
	    -p $(POSTGRES_PORT):5432 \
	    --name $(DB_CONTAINER_NAME) \
	    $(POSTGRES_IMAGE) postgres \
	    -c log_statement=all