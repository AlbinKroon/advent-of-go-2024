.DEFAULT_GOAL := vet

.PHONY:fmt vet clean
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

clean:
	go clean