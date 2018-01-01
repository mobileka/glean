# Let's make it pretty
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

# Local .env file
ifneq ("$(wildcard .env)","")
	include .env
	export $(shell sed 's/=.*//' .env)
endif

# Vars
BUILD_DIR ?= build/
NAME=glean
REPO=github.com/mobileka/${NAME}
SRC_DIRS=cmd
BINARY=glean
BINARY_SRC=$(REPO)/cmd/${NAME}

# Rules

.PHONY: all all-reaction clean glide build install all-done-reaction all-dev-reaction

all: all-reaction clean glide build install all-done-reaction

all-dev: all-dev-reaction clean glide-dev build install all-done-reaction

install:
	@printf "$(WARN_COLOR)😒 Installing...\n"
	go install -v $(BINARY_SRC)
	@printf "$(OK_COLOR)Done 😒\n\n"

build:
	@printf "$(WARN_COLOR)😒 Building your stupid project\n"
	@go build -o ${BUILD_DIR}/${BINARY} -ldflags="-s -w" ${BINARY_SRC}
	@printf "$(OK_COLOR)Done... 😒\n\n"

glide:
	@printf "$(WARN_COLOR)😒 Installing dependencies since 1891\n"
	@glide install
	@printf "$(OK_COLOR)Done 😒\n\n"

glide-dev:
	@printf "$(WARN_COLOR)😒 Installing dependencies (including dev)\n"
	@glide install
	@go get -u golang.org/x/tools/cmd/goimports
	@go get -u github.com/golang/lint/golint
	@printf "$(OK_COLOR)It was so fascinating... 😒\n\n"

lint:
	@echo "$(OK_COLOR)😒 Lemme lint your 💩codz"
	@golint $(SRC_DIRS)
	@goimports -l -w $(SRC_DIRS)

clean:
	@printf "$(WARN_COLOR)😏 Killing your build files with 🔥\n"
	if [ -d ${BUILD_DIR} ] ; then rm -rf ${BUILD_DIR} ; fi
	@printf "$(OK_COLOR)Done 😏\n\n"

all-reaction:
	@printf "$(WARN_COLOR)😒 I'm doing all the hard work...\n\n"

all-dev-reaction:
	@printf "$(ERROR_COLOR)No! $(WARN_COLOR)😡 Can you at least run them one by one? Lazy bastard.\n\n"

all-done-reaction:
	@printf "$(WARN_COLOR)😒 Let me pretend that I'm happy 🎉\n\n"
