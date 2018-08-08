FROM golang:latest as builder
ENV GOPATH /go
RUN mkdir -p /go/src/github.com/bruno-chavez/ancestorquotes/
ADD . /go/src/github.com/bruno-chavez/ancestorquotes/
WORKDIR /go/src/github.com/bruno-chavez/ancestorquotes/
RUN go get github.com/urfave/cli
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest as runner
ENV DOCKER true
# RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/bruno-chavez/ancestorquotes .

CMD ["./app"] 
