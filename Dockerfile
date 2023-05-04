#Build
FROM golang:1.19-alpine3.17 as build

WORKDIR /bot

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install

RUN ["go", "build", "-v", "-o", "/bin/bot"]

#Deploy

FROM alpine:3.17

WORKDIR /

COPY --from=build /bin/bot /bin/bot

USER root:root

CMD /bin/bot
