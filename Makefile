package = github.com/reynico/fibertel-stats

.PHONY: release

release:
	go get
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/fibertel-stats-linux-amd64 $(package)
	GOOS=linux GOARCH=arm GOARM=7 go build -o release/fibertel-stats-linux-armv7 $(package)
	GOOS=linux GOARCH=arm64 go build -o release/fibertel-stats-linux-armv64 $(package)
	GOOS=darwin GOARCH=amd64 go build -o release/fibertel-stats-darwin-amd64 $(package)
