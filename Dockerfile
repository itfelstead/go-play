# Straight from golang site. For dev + run

FROM golang:1.14

WORKDIR /go/src/app
COPY ./play ./play

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["/bin/bash"]
