FROM golang:1.12

EXPOSE 8081
WORKDIR /go/src/github.com/d0ku/distributed_math
COPY . .

RUN go install ./adder

CMD ["adder"]
