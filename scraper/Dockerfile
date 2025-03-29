FROM golang:latest AS build-debug

WORKDIR /app/

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download
RUN apt-get update && apt-get install -y \
    git \
    curl \
    && go install github.com/go-delve/delve/cmd/dlv@latest


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/scraper .

FROM golang AS debug
WORKDIR /app/



COPY --from=build-debug /app/scraper /app/scraper
COPY --from=build-debug /go/bin/dlv /dlv
CMD ["/dlv", "exec", "/app/scraper", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient"]



FROM golang:latest AS build

WORKDIR /app/

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/scraper .

FROM scratch AS prod

# Copy the application executable from the build image
COPY --from=build /app/scraper /scraper

EXPOSE 8080
CMD ["/scraper"]
