# Stage 1: Build Frontend
FROM node:22-alpine AS frontend-builder
WORKDIR /app
COPY web/app/package*.json ./web/app/
RUN cd web/app && npm install
COPY web/app/ ./web/app/
RUN cd web/app && npm run build

# Stage 2: Build Backend
FROM mcr.microsoft.com/devcontainers/go:1.24-bookworm AS backend-builder
WORKDIR /app
COPY go.mod ./
# RUN go mod download
COPY . .
RUN go build -o k8dclusterlife ./cmd/k8dclusterlife/main.go

# Stage 3: Runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=backend-builder /app/k8dclusterlife .
COPY --from=frontend-builder /app/web/app/dist ./web/app/dist
EXPOSE 8080
CMD ["./k8dclusterlife"]
