run-go:
	go test ./golang/util.go ./golang/$(arg)_test.go -v

rust-fmt:
	rustfmt ./rust/src/aoc$(arg).rs

run-rust: rust-fmt
	cargo test --manifest-path ./rust/Cargo.toml $(arg) -v -- --nocapture