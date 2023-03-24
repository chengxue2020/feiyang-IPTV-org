FROM golang:1.19-alpine AS build

WORKDIR /app

COPY ./Golang/go.mod ./
COPY ./Golang/go.sum ./
RUN go mod download

COPY ./Golang/*.go ./
COPY ./Golang/liveurls/*.go ./liveurls/

RUN go build -o /allinone

FROM alpine:3.14

COPY --from=build /allinone /allinone

EXPOSE 35455

CMD [ "/allinone" ]