LINTERCOMMAND=gometalinter --vendor --tests --skip=mock --deadline=1500s

packages = \
	./config \
	./database \
	./models \
    ./services \
    ./services/auth \
    ./services/common \
    ./services/usersCRUD \
    ./services/tasksCRUD \

all: test format go-build dc-build dc-up

test:
	@$(foreach package,$(packages), set -e; \
    		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

code-quality:
	$(LINTERCOMMAND) --checkstyle ./... > static-analysis.xml

code-quality-print:
	$(LINTERCOMMAND) ./...

format:
	go fmt $(packages)

go-build:
	go build -o cool_tasks

dc-build:
	docker-compose build

dc-up:
	docker-compose up
