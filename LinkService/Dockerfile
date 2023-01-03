FROM golang:1.19-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod .

RUN go mod download

RUN go mod tidy
RUN go build -o /main
EXPOSE 8081

CMD [ "/main" ]