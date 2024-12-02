FROM golang:1.22 AS build-stage
RUN groupadd -g 10000 projectipv6 && useradd -u 10000 -r -g projectipv6 projectipv6
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM debian:bullseye-slim AS build-release-stage
COPY --from=build-stage /etc/passwd /etc/passwd
COPY --from=build-stage /etc/group /etc/group
WORKDIR /
COPY --from=build-stage /app/main /main
RUN apt update && apt install -y bash procps
EXPOSE 8080
USER projectipv6
ENTRYPOINT ["/main"]