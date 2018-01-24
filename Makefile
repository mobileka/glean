# Let's make it pretty
NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

# Vars
BUILD_DIR ?= build
RES_DIR = resources
NAME=glean
REPO=github.com/mobileka/${NAME}
SRC_DIRS=cmd
BINARY=glean
BINARY_SRC=$(REPO)/cmd/${NAME}
PLATFORMS=linux/amd64 linux/386 windows/amd64 windows/386 darwin/amd64 darwin/386

tmp = $(subst /, ,$@)
os = $(word 1, $(tmp))
arch = $(word 2, $(tmp))

# Rules
.PHONY: all all-reaction clean glide build install all-done-reaction all-dev-reaction run $(PLATFORMS)

all: all-reaction clean glide build install all-done-reaction

all-dev: all-dev-reaction clean glide-dev build install all-done-reaction

$(PLATFORMS):
	@printf "$(WARN_COLOR)😒 Building $(os)-$(arch)...$(NO_COLOR)\n"
	GOOS=$(os) GOARCH=$(arch) go build  -o ${BUILD_DIR}/${BINARY}-$(os)-$(arch) -ldflags="-s -w" ${BINARY_SRC}
	@cp ${RES_DIR}/glean.yaml ${BUILD_DIR}/
	@printf "$(OK_COLOR)Done... 😒\n\n"

build: $(PLATFORMS)

glide:
	@printf "$(WARN_COLOR)😒 Installing dependencies since 1915\n"
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
	@printf "$(WARN_COLOR)😒 All done... Let me pretend that I'm happy 🎉\n\n"
