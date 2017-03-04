# Copyright 2017 The Nakama Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

LIBNAME := example
VERSION := 0.1.0-dev
BUILDDIR := build
COMMITID := $(shell git rev-parse --short HEAD 2>/dev/null || echo nosha)

GOFLAGS := -gcflags "-trimpath ${CURDIR}"
LDFLAGS := -ldflags "-X example.version=${VERSION} -X example.commitID=${COMMITID}"
PLATFORMS := linux

.PHONY: help
default help:
	@echo "Usage: make <command>\n"
	@echo "The commands are:"
	@echo "   all         Alias for '$(LIBNAME)' command."
	@echo "   $(LIBNAME)      Build a development version of the server. Runs dependent rules."
	@echo "   release     Build production release(s). Runs dependent rules."
	@echo "   test        Execute all development tests."
	@echo "   vet         Perform static error checks against the source.\n"

.PHONY: all
all: $(LIBNAME)

$(LIBNAME):
	go build -buildmode=plugin ${GOFLAGS} ${LDFLAGS} -o ${BUILDDIR}/dev/${LIBNAME} example.go

.PHONY: release
release:
	@echo "Not yet implemented"

.PHONY: test
test:
	@echo "Not yet implemented"

.PHONY: vet
vet:
	go vet ${GOFLAGS} ${LDFLAGS}
