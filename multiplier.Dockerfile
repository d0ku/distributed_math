FROM golang:1.12

EXPOSE 8083
WORKDIR /go/src/github.com/d0ku/distributed_math
COPY . .

RUN go install ./multiplier

CMD ["multiplier"]
