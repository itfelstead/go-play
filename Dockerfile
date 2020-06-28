# Straight from golang site. For dev + run
#
# e.g. you might run this from the root of the source dir of interest;
#   docker run -it --rm -v ${PWD}:/go/src/app --name gotest golangdev
#
FROM golang:1.14

WORKDIR /go/src/app

RUN go get -d -v ./...
RUN go install -v ./...

# For now I'm using the VS Code go extension via 'Remote-Containers: Attach to running container..'
# and that requires a few go tools to be present on the dev container;
RUN go get -u github.com/ramya-rao-a/go-outline
# TBD: gocode doesn't seem to get picked up by VS Code; why?
RUN go get -u github.com/stamblerre/gocode 
RUN go get -u github.com/uudashr/gopkgs/cmd/gopkgs
RUN go get -u github.com/rogpeppe/godef
RUN go get -u github.com/sqs/goreturns
RUN go get -u golang.org/x/lint/golint

CMD ["/bin/bash"]
