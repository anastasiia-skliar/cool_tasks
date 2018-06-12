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
create-table:
psql -h 127.0.0.1 -p 5432 -U postgres -d cool_tasks -a -f cool_tasks.sql
