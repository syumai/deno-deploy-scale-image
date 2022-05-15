.PHONY: generate
generate:
	go generate

.PHONY: run
run:
	deployctl deploy --project=scale-image ./mod.js
