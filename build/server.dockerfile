FROM golang:1.19-alpine as build

WORKDIR /groupware

COPY . .

FROM build as release




