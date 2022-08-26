FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init
RUN go build
CMD [ "./backend-core" ]
EXPOSE 8080
