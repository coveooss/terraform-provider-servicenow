SOURCES = $(wildcard **/*.go)

.PHONY: test
test:
	go test ./...

terraform-provider-servicenow: $(SOURCES)
	go build ./...

.PHONY: build
build: terraform-provider-servicenow

.PHONY: install
install: terraform-provider-servicenow
	mv terraform-provider-servicenow $(shell dirname $(shell which terraform))
