FROM node:20 as frontend

WORKDIR /app
COPY ./vue-ui/*.json ./
RUN rm ./package-lock.json
COPY ./vue-ui/*.ts ./
COPY ./vue-ui/*.html ./
RUN ls -la
COPY ./vue-ui/src ./src
COPY ./vue-ui/public ./public

RUN npm install
RUN npm run build

FROM golang:latest as backend

WORKDIR /app
RUN mkdir -p ./static/assets
RUN mkdir -p ./assets
COPY ./games-go/go.mod ./
COPY ./games-go/go.sum ./
COPY ./games-go/.env ./
COPY ./games-go/cmd ./cmd
COPY ./games-go/internal ./internal
# COPY ./games-go/server ./server
COPY --from=frontend /app/dist/index.html ./templates/index.html
COPY --from=frontend /app/dist ./ui

RUN go mod download

COPY ./vue-ui/dist/assets /assets
# COPY ./vue-ui/dist/index.html /template/index.html
COPY ./vue-ui/src/output.css /assets
# COPY ./games-go-blueprint/.env ./

RUN CGO_ENABLED=0 GOOS=linux go build -o games-kevin-fechner ./cmd/api/main.go
# RUN CGO_ENABLED=1 GOOS=linux go build -o games-kevin-fechner ./cmd/api/main.go


FROM alpine:latest

WORKDIR /app
RUN mkdir ./assets
# Copy the binary to the production image from the builder stage.
COPY --from=backend /app/games-kevin-fechner .
COPY --from=backend /app/assets .
# COPY --from=backend /app/internal ./internal
COPY --from=backend /app/ui/assets ./assets
COPY --from=backend /app/ui/index.html ./templates/index.html
COPY --from=backend /app/.env ./.env
# FROM alpine:latest

# RUN apk update
# RUN apk add curl bash nvim


EXPOSE 8080
CMD ["./games-kevin-fechner"]