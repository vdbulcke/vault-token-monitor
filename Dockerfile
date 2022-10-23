FROM alpine:latest  
RUN apk --no-cache add ca-certificates
## non privileged user
USER 1111 
# EXPOSE 9000
WORKDIR /app/
COPY vault-token-monitor /app/vault-token-monitor

ENTRYPOINT ["/app/vault-token-monitor", "server" ,"--config",  "/app/config.yaml"]