FROM golang:1.16.5 as builder
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY /server .
# Build app
CMD cd /server
RUN go build -o smsOutBound

FROM alpine:3.14 as production

WORKDIR /github.com/ddld93/sms-outboubd-microservices/server
# Copy built binary from builder
COPY --from=builder app .
# Expose port
EXPOSE 3333
# Exec built binary
CMD ./smsOutBound