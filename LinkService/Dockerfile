FROM golang:1.19-alpine
#Fixes the issue with missing c compiler 
RUN apk add build-base
ENV GOOS=linux GOARCH=amd64 

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod .

RUN go mod download

RUN go mod tidy
RUN apk update && apk add gcc librdkafka-dev zstd-libs libsasl lz4-dev libc-dev musl-dev 

RUN go build -tags musl -ldflags '-w -extldflags "-static"' -o /main
EXPOSE 8081

CMD [ "/main" ]