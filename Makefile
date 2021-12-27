.PHONY: build-all
build-all: build-static

.PHONY: run-all
run-all: run-static


.PHONY: build
build:
	go build -v ./...

.PHONY: build-static
build-static:
	rustup toolchain install nightly
	cd pkg/download && rustup run nightly cargo build --release 
	cp pkg/download/target/release/libydl.a pkg/
	go build -v ./...

.PHONY: run-static
run-static:
	RUST_LOG=trace ./ydl

# test rust lib
.PHONY: test
test:
	cd pkg/download && RUST_LOG=trace rustup run nightly cargo test --lib

# clean all packages
.PHONY: clean
clean:
	rm -rf main ydl pkg/libydl.so pkg/libydl.a pkg/download/target
