SERVER_BINARY=chunkymonkey
EXTRA_BINARIES=\
	datatests \
	inspectlevel \
	intercept \
	replay \
	style

DIAGRAMS=diagrams/top-level-architecture.png

server: $(SERVER_BINARY)

all: server extra

extra: $(EXTRA_BINARIES)

# godag leaves garbage behind.
cleantmp:
	@-rm -rf src/tmp*

clean:
	@-rm -f $(SERVER_BINARY) $(EXTRA_BINARIES)
	@-rm -rf test_obj
	@gd -q -c

fmt:
	@gd -q -fmt --tab src

check: style
	@./style `find . -name \*.go`

test: cleantmp
	@-rm -rf test_obj/tmp*
	@mkdir -p test_obj
	@gd -q -L test_obj -t src

test_data: datatests
	@./datatests

bench: cleantmp
	@-rm -rf test_obj/tmp*
	@mkdir -p test_obj
	@gd -q -L test_obj -b Benchmark -t src

# Note that this will also compile code in the src/util directory.
libs: cleantmp
	@gd -q src

chunkymonkey: libs
	@gd -q -I src -o $@ src/main

datatests: libs
	@gd -q -I src -o $@ src/util/$@

intercept: libs
	@gd -q -I src -o $@ src/util/$@

inspectlevel: libs
	@gd -q -I src -o $@ src/util/$@

replay: libs
	@gd -q -I src -o $@ src/util/$@

style: src/util/style/style.go
	@gd -q -I src -o $@ src/util/$@

docs: $(DIAGRAMS)

%.png: %.dot
	@dot -Tpng $< -o $@

.PHONY: all bench check clean cleantmp docs extra fmt libs server test test_data
