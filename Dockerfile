# Straight from golang site. For dev + run
#
# e.g. you might run this from the root of the source dir of interest;
#   docker run -it --rm -v ${PWD}:/go/src/app --name gotest golangdev
#
FROM golang:1.14

WORKDIR /go/src/app

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["/bin/bash"]
