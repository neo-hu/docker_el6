.PHONY: build

GO := go

PREFIX := $(DESTDIR)/usr/local
BINDIR := $(PREFIX)/sbin

all: containerd-build dockerd-build shim-build runc-build

runc-build:
	$(GO) build -buildmode=pie $(EXTRA_FLAGS) -o runc ./cmd/runc

shim-build:
	$(GO) build -buildmode=pie $(EXTRA_FLAGS) -o containerd-shim ./cmd/containerd-shim

dockerd-build:
	$(GO) build -buildmode=pie $(EXTRA_FLAGS) -tags "exclude_graphdriver_btrfs" -o dockerd ./cmd/dockerd

containerd-build:
	$(GO) build -buildmode=pie $(EXTRA_FLAGS) -tags "no_btrfs" -o docker-containerd ./cmd/containerd

clean:
	rm -f runc runc-* containerd-shim docker-containerd dockerd

install: install-containerd install-dockerd install-shim install-runc

install-runc:
	install -D -m0755 runc $(BINDIR)/runc

install-containerd:
	install -D -m0755 docker-containerd $(BINDIR)/docker-containerd

install-dockerd:
	install -D -m0755 dockerd $(BINDIR)/dockerd

install-shim:
	install -D -m0755 containerd-shim $(BINDIR)/containerd-shim

uninstall-runc:
	rm -f $(BINDIR)/runc

uninstall-containerd:
	rm -f $(BINDIR)/docker-containerd

uninstall-dockerd:
	rm -f $(BINDIR)/dockerd

uninstall-shim:
	rm -f $(BINDIR)/shim

uninstall: uninstall-shim uninstall-dockerd uninstall-containerd uninstall-runc


