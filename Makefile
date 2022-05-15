.PHONY: generate
generate:
	go generate

.PHONY: run
run:
	deno run --allow-net="0.0.0.0:8000,github.com,raw.githubusercontent.com" --allow-env=DENO_REGION ./mod.js
