FROM golang:1.20 AS build
WORKDIR /app
COPY . ./

RUN go build -o app cmd/api/main.go

FROM alpine:latest
COPY --from=build /app/app /app
EXPOSE 8080
CMD [ "/app" ]

