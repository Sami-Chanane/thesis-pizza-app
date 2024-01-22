FROM golang:1.22rc1-alpine3.19 AS base
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./


FROM base AS test-stage
RUN go test -v ./... -coverprofile=coverage.out > report.txt


FROM scratch AS test-report
COPY --from=test-stage /app/report.txt report.txt
COPY --from=test-stage /app/coverage.out coverage.out


FROM base AS build-stage
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping


FROM alpine:3.19.0 AS final-stage
COPY --from=build-stage /docker-gs-ping /docker-gs-ping
COPY --from=build-stage /app/static /static
COPY --from=build-stage /app/templates /templates
CMD [ "./docker-gs-ping" ]