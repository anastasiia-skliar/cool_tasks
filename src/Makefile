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

all: test format migrate-down migrate-up go-build dc-build dc-up

setup:
	go get -u gopkg.in/alecthomas/gometalinter.v2

test:
	@$(foreach package,$(packages), set -e; \
    		go test -coverprofile $(package)/cover.out -covermode=count $(package);)

perf-test:
	vegeta attack -targets=src/performance/test-login.txt -duration=5s -rate=150 | tee src/performance/results.bin | vegeta report
	cat src/performance/results.bin | vegeta report -reporter=plot > src/performance/plot.html

code-quality:
	$(LINTERCOMMAND) --checkstyle ./... > static-analysis.xml

code-quality-print:
	$(LINTERCOMMAND) ./...

format:
	go fmt $(packages)
	
migrate-down:
	goose -dir ./migrations postgres "user=postgres password=postgres host=172.19.0.2 dbname=cool_tasks sslmode=disable" down-to 0

migrate-up:
	goose -dir ./migrations postgres "user=postgres password=postgres host=172.19.0.2 dbname=cool_tasks sslmode=disable" up

doc:
	godoc -http=:6060

go-build:
	GOOS=linux GOARCH=amd64 go build -o cool_tasks

dc-build:
	docker-compose build

dc-up:
	docker-compose up
