# Copyright 2024 The tk9.0-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean edit editor test work w65 lib_win lib_linux lib_darwin lib_freebsd build_all_targets

TAR = tcl-core9.0.0rc1-src.tar.gz
URL = http://kumisystems.dl.sourceforge.net/project/tcl/Tcl/9.0.0/$(TAR)
TAR2 = tk9.0.0rc1-src.tar.gz
URL2 = http://deac-riga.dl.sourceforge.net/project/tcl/Tcl/9.0.0/$(TAR2)
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
		--exclude .git/ \
		--exclude \*.gz \
		--exclude html/ \
		.  \
		win65:src/modernc.org/tk9.0

# lib_win: download
# 	if [ "$(GOOS)" != "windows" ]; then exit 1 ; fi
# 	rm -rf ~/tmp/tcl9* ~/tmp/tk9*
# 	tar xf $(TAR) -C ~/tmp
# 	tar xf $(TAR2) -C ~/tmp
# 	sh -c "cd ~/tmp/tcl9.0.0/win ; ./configure"
# 	make -C ~/tmp/tcl9.0.0/win
# 	cp -v ~/tmp/tcl9.0.0/win/libtommath.dll ~/tmp/tcl9.0.0/win/tcl90.dll .
# 	sh -c "cd ~/tmp/tk9.0.0/win ; ./configure --with-tcl=$$HOME/tmp/tcl9.0.0/win"
# 	make -C ~/tmp/tk9.0.0/win
# 	cp -v ~/tmp/tk9.0.0/win/tcl9tk90.dll .
# 	rm -rf embed_windows/
# 	mkdir embed_windows
# 	cp ../libtk9.0/library/library.zip .
# 	unzip library.zip
# 	rm library.zip
# 	mv library/ tk_library/
# 	zip -r embed_windows/tk_library.zip tk_library/
# 	rm -rf mkdir embed_windows_$(GOARCH)
# 	mkdir embed_windows_$(GOARCH)
# 	rm -f embed_windows_$(GOARCH)/dll.zip
# 	zip embed_windows_$(GOARCH)/dll.zip *.dll
# 	rm -f *.dll

lib_win: download
	if [ "$(GOOS)" != "windows" ]; then exit 1 ; fi
	rm -rf ~/tmp/tcl9* ~/tmp/tk9* embed/$(GOOS)/$(GOARCH)
	mkdir -p embed/$(GOOS)/$(GOARCH)
	tar xf $(TAR) -C ~/tmp
	tar xf $(TAR2) -C ~/tmp
	sh -c "cd ~/tmp/tcl9.0.0/win ; ./configure"
	make -C ~/tmp/tcl9.0.0/win
	cp -v ~/tmp/tcl9.0.0/win/libtommath.dll ~/tmp/tcl9.0.0/win/tcl90.dll embed/$(GOOS)/$(GOARCH)
	sh -c "cd ~/tmp/tk9.0.0/win ; ./configure --with-tcl=$$HOME/tmp/tcl9.0.0/win"
	make -C ~/tmp/tk9.0.0/win
	cp -v ~/tmp/tk9.0.0/win/tcl9tk90.dll ~/tmp/tk9.0.0/win/libtk9.0.0.zip embed/$(GOOS)/$(GOARCH)
	zip -j embed/$(GOOS)/$(GOARCH)/lib.zip.tmp embed/$(GOOS)/$(GOARCH)/*.dll embed/$(GOOS)/$(GOARCH)/*.zip
	rm -f embed/$(GOOS)/$(GOARCH)/*.dll embed/$(GOOS)/$(GOARCH)/*.zip
	mv embed/$(GOOS)/$(GOARCH)/lib.zip.tmp embed/$(GOOS)/$(GOARCH)/lib.zip

lib_linux: download
	if [ "$(GOOS)" != "linux" ]; then exit 1 ; fi
	rm -rf ~/tmp/tcl9* ~/tmp/tk9* embed/$(GOOS)/$(GOARCH)
	mkdir -p embed/$(GOOS)/$(GOARCH)
	tar xf $(TAR) -C ~/tmp
	tar xf $(TAR2) -C ~/tmp
	sh -c "cd ~/tmp/tcl9.0.0/unix ; ./configure --disable-dll-unloading"
	make -C ~/tmp/tcl9.0.0/unix -j2
	cp -v ~/tmp/tcl9.0.0/unix/libtcl9.0.so embed/$(GOOS)/$(GOARCH)
	sh -c "cd ~/tmp/tk9.0.0/unix ; ./configure --with-tcl=$$HOME/tmp/tcl9.0.0/unix"
	make -C ~/tmp/tk9.0.0/unix -j2
	cp -v ~/tmp/tk9.0.0/unix/libtcl9tk9.0.so ~/tmp/tk9.0.0/unix/libtk9.0.0.zip embed/$(GOOS)/$(GOARCH)
	zip -j embed/$(GOOS)/$(GOARCH)/lib.zip.tmp embed/$(GOOS)/$(GOARCH)/*.so embed/$(GOOS)/$(GOARCH)/*.zip
	rm -f embed/$(GOOS)/$(GOARCH)/*.so embed/$(GOOS)/$(GOARCH)/*.zip
	mv embed/$(GOOS)/$(GOARCH)/lib.zip.tmp embed/$(GOOS)/$(GOARCH)/lib.zip

lib_darwin: download
	if [ "$(GOOS)" != "darwin" ]; then exit 1 ; fi
	rm -rf ~/tmp/tcl9* ~/tmp/tk9* embed/$(GOOS)/$(GOARCH)
	mkdir -p embed/$(GOOS)/$(GOARCH)
	tar xf $(TAR) -C ~/tmp
	tar xf $(TAR2) -C ~/tmp
	sh -c "cd ~/tmp/tcl9.0.0/unix ; ./configure"
	make -C ~/tmp/tcl9.0.0/unix -j2
	cp -v ~/tmp/tcl9.0.0/unix/libtcl9.0.dylib embed/$(GOOS)/$(GOARCH)
	sh -c "cd ~/tmp/tk9.0.0/unix ; ./configure --with-tcl=$$HOME/tmp/tcl9.0.0/unix --enable-aqua"
	make -C ~/tmp/tk9.0.0/unix -j2
	cp -v ~/tmp/tk9.0.0/unix/libtcl9tk9.0.dylib ~/tmp/tk9.0.0/unix/libtk9.0.0.zip embed/$(GOOS)/$(GOARCH)
	zip -j embed/$(GOOS)/$(GOARCH)/lib.zip.tmp embed/$(GOOS)/$(GOARCH)/*.dylib embed/$(GOOS)/$(GOARCH)/*.zip
	rm -f embed/$(GOOS)/$(GOARCH)/*.dylib embed/$(GOOS)/$(GOARCH)/*.zip
	mv embed/$(GOOS)/$(GOARCH)/lib.zip.tmp embed/$(GOOS)/$(GOARCH)/lib.zip

# use gmake
lib_freebsd: download
	if [ "$(GOOS)" != "freebsd" ]; then exit 1 ; fi
	rm -rf ~/tmp/tcl9* ~/tmp/tk9* embed/$(GOOS)/$(GOARCH)
	mkdir -p embed/$(GOOS)/$(GOARCH)
	tar xf $(TAR) -C ~/tmp
	tar xf $(TAR2) -C ~/tmp
	sh -c "cd ~/tmp/tcl9.0.0/unix ; ./configure --disable-dll-unloading"
	gmake -C ~/tmp/tcl9.0.0/unix -j2
	cp -v ~/tmp/tcl9.0.0/unix/libtcl9.0.so embed/$(GOOS)/$(GOARCH)
	sh -c "cd ~/tmp/tk9.0.0/unix ; ./configure --with-tcl=$$HOME/tmp/tcl9.0.0/unix"
	gmake -C ~/tmp/tk9.0.0/unix -j2
	cp -v ~/tmp/tk9.0.0/unix/libtcl9tk9.0.so ~/tmp/tk9.0.0/unix/libtk9.0.0.zip embed/$(GOOS)/$(GOARCH)
	zip -j embed/$(GOOS)/$(GOARCH)/lib.zip.tmp embed/$(GOOS)/$(GOARCH)/*.so embed/$(GOOS)/$(GOARCH)/*.zip
	rm -f embed/$(GOOS)/$(GOARCH)/*.so embed/$(GOOS)/$(GOARCH)/*.zip
	mv embed/$(GOOS)/$(GOARCH)/lib.zip.tmp embed/$(GOOS)/$(GOARCH)/lib.zip
