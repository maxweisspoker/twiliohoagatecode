FROM golang:1-alpine AS build

WORKDIR /app
COPY . /app/
RUN cd /app/ && go mod download && go mod tidy
RUN cd /app/ && env CGO_ENABLED=0 GOOS=linux go build -ldflags="-extldflags=-static -s -w" -o twiliohoagatecode /app

FROM scratch
COPY --from=build /app/twiliohoagatecode /twiliohoagatecode
EXPOSE 8080
ENTRYPOINT ["/twiliohoagatecode"]
