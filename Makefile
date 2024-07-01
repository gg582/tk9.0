# Copyright 2024 The tk9.0-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean edit editor test work

all: editor
	golint 2>&1
	staticcheck 2>&1

build_all_targets:
	./build_all_targets.sh
	echo done

clean:
	rm -f log-* cpu.test mem.test *.out go.work*
	go clean

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile go.mod builder.json *.go & fi

editor:
	# go test -c -o /dev/null 2>&1 | tee -a log-editor
	go build -v  -o /dev/null generator.go 2>&1 | tee -a log-editor
	go run generator.go 2>&1 | tee log-generator
	gofmt -l -s -w . 2>&1 | tee log-editor
	go build -v  -o /dev/null ./... 2>&1 | tee -a log-editor
	go test  2>&1 | tee -a log-editor

test:
	go test -v -timeout 24h -count=1 2>&1 | tee log-test
	grep -a 'TRC\|TODO\|ERRORF\|FAIL' log-test || true 2>&1 | tee -a log-test

work:
	rm -f go.work*
	go work init
	go work use .
	go work use ../cc/v4
	go work use ../ccgo/v3
	go work use ../ccgo/v4
	go work use ../libX11
	go work use ../libXft
	go work use ../libadvapi32
	go work use ../libc
	go work use ../libcomctl32
	go work use ../libcomdlg32
	go work use ../libfontconfig
	go work use ../libgdi32
	go work use ../libimm32
	go work use ../libkernel32
	go work use ../libole32
	go work use ../liboleaut32
	go work use ../libshell32
	go work use ../libtcl9.0
	go work use ../libtk9.0
	go work use ../libuser32
	go work use ../libuserenv
	go work use ../libwinspool
	go work use ../libz
	go work use ../ngrab
	go work use ../tcl9.0
