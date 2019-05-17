.PHONY: build

GO := go

SOURCES := $(shell find . 2>&1 | grep -E '.*\.(c|h|go)$$')
PREFIX := $(DESTDIR)/usr/local
BINDIR := $(PREFIX)/sbin


runc: $(SOURCES)
	$(GO) build -buildmode=pie $(EXTRA_FLAGS) -o runc .

all: runc

clean:
	rm -f runc runc-*

_allpackages = $(shell $(GO) list ./... | grep -v vendor)
allpackages = $(if $(__allpackages),,$(eval __allpackages := $$(_allpackages)))$(__allpackages)