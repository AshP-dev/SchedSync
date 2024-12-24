FROM golang:1.23-alpine AS build

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Enable CGO and build the binary
ENV CGO_ENABLED=1
RUN go build -o server ./main.go

FROM alpine
WORKDIR /app
COPY --from=build /app/server /app
EXPOSE 3010
CMD ["/app/server"]