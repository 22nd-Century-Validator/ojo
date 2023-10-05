#!/usr/bin/make -f


BRANCH         := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT         := $(shell git log -1 --format='%H')
BUILD_DIR      ?= $(CURDIR)/build
DIST_DIR       ?= $(CURDIR)/dist
LEDGER_ENABLED ?= true
TM_VERSION     := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')
DOCKER         := $(shell which docker)
PROJECT_NAME   := ojo
HTTPS_GIT      := https://github.com/ojo-network/ojo.git

###############################################################################
##                                  Version                                  ##
###############################################################################

ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

###############################################################################
##                                   Build                                   ##
###############################################################################

build_tags = netgo

ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=ojo \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=ojod \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/cometbft/cometbft/version.TMCoreSemVer=$(TM_VERSION)

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

build: go.sum
	@echo "--> Building..."
	go build -mod=readonly $(BUILD_FLAGS) -o $(BUILD_DIR)/ ./...

build-no_cgo:
	@echo "--> Building static binary with no CGO nor GLIBC dynamic linking..."
	CGO_ENABLED=0 CGO_LDFLAGS="-static" $(MAKE) build

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

install: go.sum
	@echo "--> Installing..."
	go install -mod=readonly $(BUILD_FLAGS) ./...

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

clean:
	@echo "--> Cleaning..."
	@rm -rf $(BUILD_DIR)/**  $(DIST_DIR)/**

.PHONY: install build build-linux clean

###############################################################################
##                                  Docker                                   ##
###############################################################################

docker-build:
	@docker build -t ojo .

.PHONY: .docker-build

###############################################################################
##                                   Tests                                   ##
###############################################################################

PACKAGES_UNIT=$(shell go list ./... | grep -v -e '/tests/e2e' -e 'tests/simulation')
PACKAGES_E2E=$(shell go list ./... | grep '/e2e')
TEST_PACKAGES=./...
TEST_TARGETS := test-unit test-unit-cover test-race
TEST_COVERAGE_PROFILE=coverage.txt

test-unit: ARGS=-timeout=10m -tags='norace'
test-unit: TEST_PACKAGES=$(PACKAGES_UNIT)
test-unit-cover: ARGS=-timeout=10m -tags='norace' -coverprofile=$(TEST_COVERAGE_PROFILE) -covermode=atomic
test-unit-cover: TEST_PACKAGES=$(PACKAGES_UNIT)
test-race: ARGS=-timeout=10m -race
test-race: TEST_PACKAGES=$(PACKAGES_UNIT)
$(TEST_TARGETS): run-tests

run-tests:
ifneq (,$(shell which tparse 2>/dev/null))
	@echo "--> Running tests"
	@go test -mod=readonly -json $(ARGS) $(TEST_PACKAGES) | tparse
else
	@echo "--> Running tests"
	@go test -mod=readonly $(ARGS) $(TEST_PACKAGES)
endif

test-e2e:
	$(MAKE) docker-build
	@go test -mod=readonly -race $(PACKAGES_E2E) -v

cover-html: test-unit-cover
	@echo "--> Opening in the browser"
	@go tool cover -html=$(TEST_COVERAGE_PROFILE)

.PHONY: cover-html run-tests test-e2e $(TEST_TARGETS)

###############################################################################
###                                Linting                                  ###
###############################################################################

golangci_lint_cmd=golangci-lint

lint:
	@echo "--> Running linter with revive"
	@go install github.com/mgechev/revive
	@revive -config .revive.toml -formatter friendly ./...

lint-fix:
	@echo "--> Running linter to fix the lint issues"
	@go install mvdan.cc/gofumpt
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint
	@$(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0 --timeout=8m
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name "*.pb.go" -not -name "*.pb.gw.go" -not -name "*.pulsar.go" -not -path "./crypto/keys/secp256k1/*" | xargs gofumpt -w -l

.PHONY: lint lint-fix

###############################################################################
##                                Simulations                                ##
###############################################################################

SIMAPP = ./tests/simulation

test-sim-non-determinism:
	@echo "Running non-determinism application simulations..."
	@go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=200 -Commit=true -Period=0 -v -timeout 24h

test-sim-multi-seed-short:
	@echo "Running short multi-seed application simulations. This may take a while!"
	@runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 10 TestFullAppSimulation

test-sim-benchmark-invariants:
	@echo "Running simulation invariant benchmarks..."
	@go test -mod=readonly $(SIMAPP) -benchmem -bench=BenchmarkFullAppSimulation -run=NOOP \
	-Enabled=true -NumBlocks=1000 -BlockSize=200 -Period=1 -Commit=true -Seed=57 -v -timeout 24h

.PHONY: \
test-sim-non-determinism \
test-sim-multi-seed-short \
test-sim-benchmark-invariants

###############################################################################
##                                 Protobuf                                  ##
###############################################################################

protoVer=0.11.6
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-format proto-lint proto-gen
.PHONY: proto-all proto-gen proto-lint proto-check-breaking proto-format proto-swagger-gen

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/proto-gen.sh

proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main

###############################################################################
##                                Simulations                                ##
###############################################################################

SIMAPP = ./tests/simulation

# Install the runsim binary
runsim: $(RUNSIM)
$(RUNSIM):
	@echo "Installing runsim..."
	@go install github.com/cosmos/tools/cmd/runsim@v1.0.0

test-app-non-determinism:
	@echo "Running non-determinism application simulations..."
	@go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=200 -Commit=true -Period=0 -v -timeout 24h

test-app-multi-seed-short: runsim
	@echo "Running short multi-seed application simulations. This may take a while!"
	@runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 10 TestFullAppSimulation

test-app-multi-seed-long: runsim
	@echo "Running long multi-seed application simulation. This may take awhile!"
	@runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 500 50 TestFullAppSimulation

test-app-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 5 TestAppImportExport

test-app-after-import: runsim
	@echo "Running application simulation-after-import. This may take several minutes..."
	@runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 5 TestAppSimulationAfterImport

test-app-benchmark-invariants:
	@echo "Running simulation invariant benchmarks..."
	@go test -mod=readonly $(SIMAPP) -benchmem -bench=BenchmarkFullAppSimulation -run=NOOP \
	-Enabled=true -NumBlocks=1000 -BlockSize=200 -Period=1 -Commit=true -Seed=57 -v -timeout 24h

# TODO: remove
start:
	${MAKE} build
	./scripts/single-node.sh
	sleep 4
	./develop.sh

.PHONY: start