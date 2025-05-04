FROM golang:latest AS tester

WORKDIR /app

COPY ../go.mod ../go.sum ./

COPY ./db .

RUN go mod tidy

#TODO ignore buildfiles

CMD ["go", "run", "main.go"]