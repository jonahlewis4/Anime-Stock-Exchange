FROM golang:latest as build

WORKDIR /app/

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/scraper .

FROM scratch as prod

# Copy the application executable from the build image
COPY --from=build /app/scraper /scraper

EXPOSE 8080
CMD ["/scraper"]
