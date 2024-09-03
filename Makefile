# Copyright 2024 The tk9.0-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean edit editor test work w65 dlls build_all_targets

TAR = tcl9.0b3-src.tar.gz
URL = http://prdownloads.sourceforge.net/tcl/$(TAR)
TAR2 = tk9.0b3-src.tar.gz
URL2 = http://prdownloads.sourceforge.net/tcl/$(TAR2)
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)

build_all_targets:
	GOOS=darwin GOARCH=arm64 go build
	GOOS=darwin GOARCH=arm64 go test -o /dev/null -c
	GOOS=linux GOARCH=386 go build
	GOOS=linux GOARCH=386 go test -o /dev/null -c
	GOOS=linux GOARCH=amd64 go build
	GOOS=linux GOARCH=amd64 go test -o /dev/null -c
	GOOS=linux GOARCH=arm go build
	GOOS=linux GOARCH=arm go test -o /dev/null -c
	GOOS=linux GOARCH=arm64 go build
	GOOS=linux GOARCH=arm64 go test -o /dev/null -c
	GOOS=linux GOARCH=loong64 go build
	GOOS=linux GOARCH=loong64 go test -o /dev/null -c
	GOOS=linux GOARCH=ppc64le go build
	GOOS=linux GOARCH=ppc64le go test -o /dev/null -c
	GOOS=linux GOARCH=riscv64 go build
	GOOS=linux GOARCH=riscv64 go test -o /dev/null -c
	GOOS=linux GOARCH=s390x go build
	GOOS=linux GOARCH=s390x go test -o /dev/null -c
	GOOS=windows GOARCH=386 go build
	GOOS=windows GOARCH=386 go test -o /dev/null -c
	GOOS=windows GOARCH=amd64 go build
	GOOS=windows GOARCH=amd64 go test -o /dev/null -c

all: editor
	golint 2>&1
	staticcheck 2>&1

clean:
	rm -f log-* cpu.test mem.test *.out go.work*
	go clean

download:
	@if [ ! -f $(TAR) ]; then wget $(URL) ; fi
	@if [ ! -f $(TAR2) ]; then wget $(URL2) ; fi

edit:
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile go.mod builder.json *.go & fi

editor:
	go test -c -o /dev/null
	go build -v  -o /dev/null generator.go
	go run generator.go
	gofmt -l -s -w .
	go build -v  -o /dev/null
	$(shell for f in _examples/*.go ; do go build -o /dev/null $$f ; done)

test:
	go test -vet=off -v -timeout 24h -count=1

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../cc/v4
	go work use ../ccgo/v3
	go work use ../ccgo/v4
	go work use ../libc
	go work use ../libtcl9.0
	go work use ../libtk9.0
	go work use ../libz
	go work use ../tcl9.0
	go work use ../ngrab

win65:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go test -o /dev/null -c
	rsync \
		-avP \
		--rsync-path='wsl rsync' \
		--exclude \*.gz \
		--exclude .git/ \
		.  \
		win65:src/modernc.org/tk9.0

dlls: download
	if [ "$(GOOS)" != "windows" ]; then exit 1 ; fi
	rm -rf ~/tmp/tcl9* ~/tmp/tk9*
	tar xf tcl9.0b3-src.tar.gz -C ~/tmp
	tar xf tk9.0b3-src.tar.gz -C ~/tmp
	sh -c "cd ~/tmp/tcl9.0b3/win ; ./configure"
	make -C ~/tmp/tcl9.0b3/win
	cp -v ~/tmp/tcl9.0b3/win/libtommath.dll ~/tmp/tcl9.0b3/win/tcl90.dll .
	sh -c "cd ~/tmp/tk9.0b3/win ; ./configure --with-tcl=$$HOME/tmp/tcl9.0b3/win"
	make -C ~/tmp/tk9.0b3/win
	cp -v ~/tmp/tk9.0b3/win/tcl9tk90.dll .
	rm -rf embed_windows/
	mkdir embed_windows
	cp ../libtk9.0/library/library.zip .
	unzip library.zip
	rm library.zip
	mv library/ tk_library/
	zip -r embed_windows/tk_library.zip tk_library/
	rm -rf mkdir embed_windows_$(GOARCH)
	mkdir embed_windows_$(GOARCH)
	rm -f embed_windows_$(GOARCH)/dll.zip
	zip embed_windows_$(GOARCH)/dll.zip *.dll
	rm -f *.dll
