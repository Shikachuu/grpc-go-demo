default:
    @just --list

# download and install binaries for development
bootstrap: _install-protoc _install-grpcurl

# build server binary
build:
    @go build -a -o bin/templates-srv cmd/server.go

# run feature tests
test:
    @go test -cover -v ./...

# generate the gRPC client/server stubs
grpc: _install-protoc
    @hack/bin/protoc --go-grpc_out=. --go_out=. --go_opt=paths=source_relative proto/templates.proto

_install-protoc:
    #!/usr/bin/env bash
    if [[ ! -f hack/bin/protoc ]]; then
        ARCH="{{ arch() }}"
        [[ "$ARCH" == "aarch64" ]] && ARCH="aarch_64"
        OS="{{ os() }}"
        [[ "$OS" == "macos" ]] && OS="osx"

        curl -sSL "https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-$OS-$ARCH.zip" -o protoc.zip
        unzip protoc.zip bin/protoc -d hack
        chmod +x hack/bin/protoc
        rm protoc.zip
    fi

_install-grpcurl:
    #!/usr/bin/env bash
    if [[ ! -f hack/bin/grpcurl ]]; then
        VERSION="1.9.1"
        ARCH="{{ arch() }}"
        [[ "$ARCH" == "aarch64" ]] && ARCH="arm64"
        OS="{{ os() }}"
        [[ "$OS" == "macos" ]] && OS="osx"
        curl -sSL "https://github.com/fullstorydev/grpcurl/releases/download/v$VERSION/grpcurl_${VERSION}_${OS}_${ARCH}.tar.gz" | tar -xz -C hack/bin grpcurl
        chmod +x hack/bin/grpcurl
    fi