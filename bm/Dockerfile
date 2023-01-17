FROM golang:1.14-alpine as build

#declare
# ENV serverPort ${serverPort:-5000}
WORKDIR /src 
RUN apk add --no-cache git curl

#copy sources
COPY src/main.go ./ 
COPY src/.env ./ 

#copy dependencies
COPY go.mod ./
COPY go.sum ./

#install dependencies
RUN go mod download

RUN go build -o bookmaker

FROM alpine:3.9 as runtime

COPY --from=build /src/bookmaker /app/bookmaker
# EXPOSE ${serverPort}
ENTRYPOINT [ "/app/bookmaker" ]