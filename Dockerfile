FROM golang:latest as build

WORKDIR /app
RUN mkdir -p ./static/assets
COPY ./games-go-blueprint/go.mod ./
COPY ./games-go-blueprint/go.sum ./
COPY ./games-go-blueprint/cmd ./cmd
COPY ./games-go-blueprint/internal ./internal
# COPY ./games-go-blueprint/server ./server
COPY ./vue-ui/dist /static
COPY ./vue-ui/src/output.css /static
COPY ./games-go-blueprint/.env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o games-kevin-fechner ./cmd/api/main.go

FROM alpine:latest

WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=build /app/games-kevin-fechner .
# FROM alpine:latest

# RUN apk update
# RUN apk add curl bash nvim


EXPOSE 8080
CMD ["./games-kevin-fechner"]