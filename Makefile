.PHONY: build-all
build-all: build-static

.PHONY: run-all
run-all: run-static

.PHONY: build-static
build-static:
	rustup toolchain install nightly
	cd pkg/download && cargo +nightly build --release
	cp pkg/download/target/release/libydl.a pkg/
	go build -v ./...

.PHONY: run-static
run-static:
	RUST_LOG=trace ./ydl

# test rust lib
.PHONY: test-rs
test-rs:
	cd pkg/download && RUST_LOG=trace cargo test -- --nocapture

# clean all packages
.PHONY: clean
clean:
	rm -rf main_static pkg/libydl.so pkg/libydl.a pkg/download/target
