FROM alpine:3.14 

WORKDIR /github.com/ddld93/sms-outboubd-microservices/
# Copy built binary from builder
COPY . .
# Expose port
EXPOSE 3333
# Exec built binary
CMD ./smsOutBound