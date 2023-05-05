FROM golang:1.14-alpine as build

#declare
ENV port ${port:-5000}
WORKDIR /app 

RUN apk add --no-cache git curl

#copy dependencies
COPY go.mod /app 
COPY go.sum /app 

#install dependencies
RUN go mod download

# RUN go get github.com/githubnemo/CompileDaemon
# COPY ./app/* /app
# ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main

RUN go build -o bookmaker

# ******************************************************* #

FROM alpine:3.9 as runtime

WORKDIR /app 

COPY --from=build /app/bookmaker /app/bookmaker
# COPY --from=build /app/static/ /app/static/

EXPOSE ${port}
ENTRYPOINT [ "/app/bookmaker" ]
# CMD [ "sleep" "inf" ]