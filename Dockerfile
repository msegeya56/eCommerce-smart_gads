# build stage
FROM golang:1.20.4-alpine3.18 AS build-stage

# Maintainer info
LABEL maintainer="Msegeya Eric <pedencinc@gmail.com>"

WORKDIR /home/app
COPY . .
RUN apk update && apk add --no-cache git
RUN go mod download
RUN go build -v -o /home/build/api ./cmd/api

# Final stage
FROM gcr.io/distroless/static-debian11

# Maintainer info
LABEL maintainer="Msegeya Eric <pedencinc@gmail.com>"

WORKDIR /home/app

COPY --from=build-stage /home/build/api /home/app/api
COPY --from=build-stage /home/app/views /home/app/views
COPY .env /home/app/.env

CMD ["/home/app/api"]
