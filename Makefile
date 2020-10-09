install:
	go install -v

fmt:
	go fmt
	cd ./cmd && go fmt
	cd ./core && go fmt

certs:
	openssl genrsa \
		-out ./certs/localhost.key \
		2048
	opensslreqq \
		-new -x509 \
		-key ./certs/localhost.key \
		-out ./certs/localhost.cert \
		-days 3650 \
		-subj /CN=localhost

grpc:
	cd messaging
	protoc -I messaging/ messaging/*.proto --go_out=plugins=grpc:messaging

.PHONY: fmt install grpc certs
