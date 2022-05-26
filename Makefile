BINARY=bin/user

Program=user
Version=`git rev-parse --abbrev-ref HEAD`
CommitID=`git rev-parse HEAD`
LastCommitTime=`git log -1 --format=%ct | xargs -I {} date -d @{} -u +%Y-%m-%dT%H:%M:%S%Z`
BuildTime=`date -u +%Y-%m-%dT%H:%M:%S%Z`

# Setup the -Idflags options for go build here,interpolate the variable values
LDFLAGS=-ldflags "-X main.Program=${Program} -X main.Version=${Version} -X main.CommitID=${CommitID} -X main.LastCommitTime=${LastCommitTime} -X main.BuildTime=${BuildTime}"

# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY}

run:
	go run main.go -env ./.env.dev

test:
	go test -v ./tests/...

build-docker: build
	docker build . -t restfulapi/user

run-docker: build-docker
	docker run -p 8080:8080 restfulapi/user
