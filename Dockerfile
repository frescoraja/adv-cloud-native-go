FROM golang:1.8.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git
RUN go get github.com/gin-gonic/gin

ENV SOURCES /go/src/adv-cloud-native-go
COPY . ${SOURCES}

WORKDIR ${SOURCES}
RUN CGO_ENABLED=0 go build

EXPOSE 8080

ENTRYPOINT ${SOURCES}/adv-cloud-native-go
