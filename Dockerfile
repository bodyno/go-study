FROM golang

RUN mkdir -p /go/src/go-study

WORKDIR /go/src/go-study

COPY . /go/src/go-study

RUN go get -v

RUN go build

EXPOSE 8080

RUN ./go-study

