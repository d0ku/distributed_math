FROM golang:1.12

EXPOSE 8084
WORKDIR /go/src/github.com/d0ku/distributed_math
COPY . .

RUN go install ./divider

CMD ["divider"]
