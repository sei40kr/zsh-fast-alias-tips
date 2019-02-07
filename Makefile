.PHONY: all
all:
	go build -o build/def-matcher def-matcher.go
.PHONY: test
clean:
	go test -v def-matcher.go def-matcher_test.go
.PHONY: clean
clean:
	rm -f build/def-matcher
