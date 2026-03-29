PACK := phase
ORG := jgautheron
PROJECT := github.com/$(ORG)/pulumi-$(PACK)
PROVIDER_PATH := provider
VERSION_PATH := $(PROVIDER_PATH)/pkg/version.Version
CODEGEN := pulumi-tfgen-$(PACK)
PROVIDER := pulumi-resource-$(PACK)
WORKING_DIR := $(shell pwd)

PROVIDER_VERSION ?= 0.1.0-alpha.0+dev

LDFLAGS_STRIP_SYMBOLS := -s -w
LDFLAGS_PROJ_VERSION := -X $(PROJECT)/$(VERSION_PATH)=$(PROVIDER_VERSION)
LDFLAGS := $(LDFLAGS_PROJ_VERSION) $(LDFLAGS_STRIP_SYMBOLS)

_ := $(shell mkdir -p .make bin)

.PHONY: build provider tfgen schema generate_sdks build_sdks clean help generate_go build_go generate_nodejs build_nodejs generate_python build_python

build: provider build_sdks

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build         Build provider and all SDKs"
	@echo "  provider      Build the provider binary"
	@echo "  tfgen         Generate schema from Terraform provider"
	@echo "  build_sdks    Build all language SDKs"
	@echo "  clean         Remove generated files"

# Schema generation
tfgen: schema
schema: .make/schema
.make/schema: bin/$(CODEGEN)
	$(WORKING_DIR)/bin/$(CODEGEN) schema --out provider/cmd/$(PROVIDER)
	(cd provider && VERSION=$(PROVIDER_VERSION) go generate cmd/$(PROVIDER)/main.go)
	@touch $@

bin/$(CODEGEN): provider/*.go provider/go.*
	(cd provider && go build -o $(WORKING_DIR)/bin/$(CODEGEN) -ldflags "$(LDFLAGS_PROJ_VERSION)" $(PROJECT)/$(PROVIDER_PATH)/cmd/$(CODEGEN))

# Provider binary
provider: bin/$(PROVIDER)
bin/$(PROVIDER): .make/schema
	(cd provider && go build -o $(WORKING_DIR)/bin/$(PROVIDER) -ldflags "$(LDFLAGS)" $(PROJECT)/$(PROVIDER_PATH)/cmd/$(PROVIDER))

# SDK generation
generate_sdks: generate_go generate_nodejs generate_python
build_sdks: build_go build_nodejs build_python

generate_go: .make/generate_go
build_go: .make/build_go
.make/generate_go: bin/$(CODEGEN)
	$(WORKING_DIR)/bin/$(CODEGEN) go --out sdk/go/
	@touch $@
.make/build_go: .make/generate_go
	cd sdk && go list "$$(grep -e "^module" go.mod | cut -d ' ' -f 2)/go/..." | xargs -I {} bash -c 'go build {} && go clean -i {}'
	@touch $@

generate_nodejs: .make/generate_nodejs
build_nodejs: .make/build_nodejs
.make/generate_nodejs: bin/$(CODEGEN)
	$(WORKING_DIR)/bin/$(CODEGEN) nodejs --out sdk/nodejs/
	printf "module fake_nodejs_module // Exclude this directory from Go tools\n\ngo 1.17\n" > sdk/nodejs/go.mod
	@touch $@
.make/build_nodejs: .make/generate_nodejs
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ 2>/dev/null || true
	@touch $@

generate_python: .make/generate_python
build_python: .make/build_python
.make/generate_python: bin/$(CODEGEN)
	$(WORKING_DIR)/bin/$(CODEGEN) python --out sdk/python/
	printf "module fake_python_module // Exclude this directory from Go tools\n\ngo 1.17\n" > sdk/python/go.mod
	cp README.md sdk/python/ 2>/dev/null || true
	@touch $@
.make/build_python: .make/generate_python
	cd sdk/python/ && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		rm ./bin/go.mod && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build==1.2.1 && \
		cd ./bin && \
		../venv/bin/python -m build .
	@touch $@

clean:
	rm -rf sdk/{go,nodejs,python}
	rm -rf bin/*
	rm -rf .make/*
