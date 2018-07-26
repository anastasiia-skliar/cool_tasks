LINTERCOMMAND=gometalinter --vendor --tests --skip=mock --deadline=1500s

packages = \
	./config \
	./database \
	./models \
    ./services/auth \
    ./services/users \
    ./services/tasks \
    ./services/events \
    ./services/flights \
    ./services/museums \
    ./services/restaurants \
    ./services/trains \
    ./services/events \
    ./services/hotels \

all: test format go-build dc-build dc-up

setup:
	go get -u gopkg.in/alecthomas/gometalinter.v2

test:
	@$(foreach package,$(packages), set -e; \
    		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

code-quality:
	$(LINTERCOMMAND) --checkstyle ./... > static-analysis.xml

code-quality-print:
	$(LINTERCOMMAND) ./...

format:
	go fmt $(packages)

doc:
	godoc -http=:6060

go-build:
	GOOS=linux GOARCH=amd64 go build -o cool_tasks

dc-build:
	docker-compose build

dc-up:
	docker-compose up
