FROM arm32v7/golang:1.25-rc-alpine3.21  as builder

ENV GOOS="linux"
ENV CGO_ENABLED="0"

ARG CMD="http"

WORKDIR /app

COPY cmd/${CMD} cmd/${CMD}
COPY internal internal

COPY go.mod go.mod
COPY go.sum go.sum

RUN go build -o main ./cmd/${CMD}/main.go

ENTRYPOINT  ["/app/main"]

EXPOSE 8000

# FROM arm32v7/golang:1.25-rc-alpine3.21 as prod
# COPY --from=builder /app/main /bin/
# COPY internal/handler/templates internal/handler/templates
# ENTRYPOINT  ["/bin/main"]
