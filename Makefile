
all: test test-race lint
test: test-go
lint: lint-go
	
# Go
####
GO_PACKAGES := projectname/...
test-go: vet-go
	go test -p 1 -v ${GO_PACKAGES}
test-race: test
	go test --race -v ${GO_PACKAGES}
vet-go:
	go vet ${GO_PACKAGES}
lint-go:
	golint -set_exit_status ${GO_PACKAGES}
