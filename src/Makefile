LINTERCOMMAND=gometalinter --vendor --tests --skip=mock --deadline=1500s

packages = \
	./config \
	./services/welcome

.PHONY: test
test:
	@$(foreach package,$(packages), \
		set -e; \
		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

.PHONY: code-quality
code-quality:
	$(LINTERCOMMAND) --checkstyle ./... > static-analysis.xml

.PHONY: code-quality-print
code-quality-print:
	$(LINTERCOMMAND) ./...
