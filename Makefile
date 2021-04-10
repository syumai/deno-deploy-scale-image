.PHONY: generate
generate:
	go generate

.PHONY: run
run:
	deployctl run --no-check mod.js
