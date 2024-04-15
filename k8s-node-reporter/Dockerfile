FROM cgr.dev/chainguard/go:latest

RUN mkdir /app/
COPY k8s-node-reporter/main.go /app/
RUN CGO_ENABLED=0 go build -o /bin/k8s-reporter /app/main.go

EXPOSE 8080

ENTRYPOINT ["/bin/k8s-reporter"]
