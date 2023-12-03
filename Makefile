.PHONY: all
all:
	go build -o build/def-matcher def-matcher.go
.PHONY: test
test:
	go test ./...
.PHONY: clean
clean:
	rm -f build/def-matcher
