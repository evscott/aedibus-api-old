FROM golang:latest

LABEL maintainer="Eliot Scott <eliotvscott@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o jac .

EXPOSE 8080

CMD ["./jac"]