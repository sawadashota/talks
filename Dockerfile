FROM golang:1.12

ENV GO111MODULE="off"

WORKDIR /go/src/github.com/sawadashota/talks

RUN go get golang.org/x/tools/cmd/present

EXPOSE 3999
CMD ["present", "-notes", "-http", "0.0.0.0:3999"]