build:
	@echo "Building..."
	@go build chroma.gen.go
	@echo "Done."

update-pkg-cache:
    GOPROXY=https://proxy.golang.org GO111MODULE=on \
    go get github.com/CSXL/go-chroma@v$(VERSION)
