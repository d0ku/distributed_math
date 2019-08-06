FROM golang:1.12

RUN apt-get update && \
    apt-get install ragel && \
    go get golang.org/x/tools/cmd/goyacc

EXPOSE 8080
WORKDIR /go/src/github.com/d0ku/distributed_math
COPY . .

RUN go generate ./expr
RUN go install ./expr

CMD ["expr"]
