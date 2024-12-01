run-go:
	go test ./golang/util.go ./golang/$(arg)_test.go -v

run-rust:
	cargo test --manifest-path ./rust/Cargo.toml $(arg) -v -- --nocapture