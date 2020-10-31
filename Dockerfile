FROM golang:1.15.3-alpine
COPY . /go/src/cloudstorageapi.com
WORKDIR /go/src/cloudstorageapi.com

RUN apk add git
RUN go get -d ./...

CMD ["go","run","api.go"]
EXPOSE 8000
