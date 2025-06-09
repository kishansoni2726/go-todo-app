FROM golang:alpine

WORKDIR /opt/todo-app

COPY go.sum go.mod /opt/todo-app/

RUN go mod download

COPY *.go .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo

EXPOSE 5000

CMD [ "/todo" ]