LINTERCOMMAND=gometalinter --vendor --tests --skip=mock --deadline=1500s

packages = \
	./config \
	./database \
	./model \
    ./service/users \
    ./service/tasks \
    ./service/events \
    ./service/flights \
    ./service/museums \
    ./service/trains \
    ./service/hotels \

all: setup test format go-build dc-build dc-up migrate-down migrate-up

setup:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	go get github.com/pressly/goose/cmd/goose

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

doc:
	godoc -http=:6060
	http://localhost:6060/pkg/github.com/Nastya-Kruglikova/cool_tasks/

go-build:
	GOOS=linux GOARCH=amd64 go build -o cool_tasks

dc-build:
	docker-compose build
dc-up:
	docker-compose up &

migrate-down:
	goose -dir ./migrations postgres "user=rkaillncwvmdfm password=e91abd74cbc3c335f43afcbb6a5907d073d977bc43fe62f5cb5276bb4734f2a3 host=ec2-54-247-101-202.eu-west-1.compute.amazonaws.com dbname=d23vu943n7c848 sslmode=require" down-to 0

migrate-up:
	goose -dir ./migrations postgres "user=rkaillncwvmdfm password=e91abd74cbc3c335f43afcbb6a5907d073d977bc43fe62f5cb5276bb4734f2a3 host=ec2-54-247-101-202.eu-west-1.compute.amazonaws.com dbname=d23vu943n7c848 sslmode=require" up
