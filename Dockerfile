FROM golang:1.20-bullseye AS build
# FROM cosmtrek/air

ENV port ${port:-9090}

# WORKDIR /go/src/github.com/Slvr-one/bookmaker
WORKDIR /app/

RUN apt install git curl bash

#copy dependencies
COPY go.mod . 
COPY go.sum .

# RUN go get github.com/githubnemo/CompileDaemon

COPY ./app/* ./

# Unit tests
# RUN CGO_ENABLED=0 go test -v

# ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main

# RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \  
#     && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN go mod download \
  && CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/main .
# ENTRYPOINT "sleep" "50"

# ******************************************************* #

FROM alpine:3.9 AS runtime
# FROM alpine:latest  

RUN apk add --no-cache rsyslog curl bash sudo ca-certificates
WORKDIR /root/

COPY --from=build /app/bin/main ./
# COPY --from=build /app/static/ /app/static/

EXPOSE ${port}
ENTRYPOINT ["./main"]
