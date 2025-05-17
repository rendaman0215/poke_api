BUF := buf
PROTO_PATH          := proto
PROTO_OUT_PATH      := gen
PROTO_GO_OUT_PATH   := $(PROTO_OUT_PATH)/go

.PHONY: lint
lint: ## Lint proto files
lint: $(BUF) $(PROTOLINT) $(PROTOLINT_PLUGIN)
	@$(BUF) lint $(PROTO_PATH)
	@$(PROTOLINT) -plugin $(PROTOLINT_PLUGIN) $(PROTO_PATH)

.PHONY: lintfix
lintfix: ## Lint proto files with fix options
lintfix: $(BUF) $(PROTOLINT) $(PROTOLINT_PLUGIN)
	@$(BUF) lint $(PROTO_PATH)
	@$(PROTOLINT) -fix -plugin $(PROTOLINT_PLUGIN) $(PROTO_PATH)

.PHONY: generate
generate: generate/go

.PHONY: generate/go
generate/go:
	@$(BUF) generate $(PROTO_PATH) --template buf.gen.go.yaml
