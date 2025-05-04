FROM golang:latest AS tester

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod tidy

#TODO ignore buildfiles
COPY ./db .

CMD ["go", "run", "main.go"]