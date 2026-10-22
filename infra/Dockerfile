# Build from the repo root: docker build -f infra/Dockerfile -t kuber-go .
FROM golang:1.26 AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /out/server ./cmd/api

FROM gcr.io/distroless/static-debian12

COPY --from=build /out/server /server

EXPOSE 5000

ENTRYPOINT ["/server"]
