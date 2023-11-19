FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /tw-employee-data-service ./cmd/api

EXPOSE 4000

CMD ["/tw-employee-data-service"]
