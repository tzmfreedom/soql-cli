NAME := soql
SRCS := $(shell find . -type d -name vendor -prune -o -type f -name "*.go" -print)
VERSION := 0.1.0
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\""
DIST_DIRS := find * -type d -exec
ANTLR := java -Xmx500M -cp "/usr/local/lib/antlr-4.7.1-complete.jar:$(CLASSPATH)" org.antlr.v4.Tool

.PHONY: run
run: format
	@go run . run -a "Foo#action" -f example.cls

.PHONY: test
test: format
	@go test ./...

.PHONY: build
build: format
	@go build

.PHONY: format
format: import
	@gofmt -w main.go parser/parser_test.go parser/statement_builder.go evaluator.go

.PHONY: import
import:
ifneq ($(shell command -v goimports 2> /dev/null),)
	@goimports -w .
endif

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
endif

.PHONY: deps
deps: dep
	dep ensure

.PHONY: cross-build
cross-build: # deps
	-@goimports -w $(SRCS)
	@gofmt -w $(SRCS)
	@for os in darwin linux windows; do \
	    for arch in 386 amd64; do \
	        GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -a -tags netgo \
	        -installsuffix netgo $(LDFLAGS) -o dist/$$os-$$arch/$(NAME); \
	    done; \
	done

.PHONY: dist
dist:
	@cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar zcf $(NAME)-$(VERSION)-{}.tar.gz {} \;

.PHONY: generate
generate:
	$(ANTLR) -Dlanguage=Go -visitor ./parser/dml.g4

.PHONY: tag
tag:
	git tag v$(VERSION) -f
	git push origin v$(VERSION) -f
