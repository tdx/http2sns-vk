BUILD_ARGS:=GOOS=linux GOARCH=amd64

build:
	$(BUILD_ARGS) go build -o bin/http2sns ./cmd/http2sns
	chmod +x bin/http2sns