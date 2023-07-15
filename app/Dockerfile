FROM golang:1.20-bullseye AS build
# FROM cosmtrek/air

ENV port ${port:-9090}
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# WORKDIR /go/src/github.com/Slvr-one/bookmaker
WORKDIR /app

RUN apt install git curl bash

COPY . .

# RUN go get github.com/githubnemo/CompileDaemon
# RUN go install github.com/swaggo/swag/cmd/swag@latest 
# RUN go install github.com/go-task/task/v3/cmd/task@latest
# RUN task docs


# RUN task build

# Unit tests
# RUN go test -v

# ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main

# RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \  
#     && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN go mod download \
  && go build -a -installsuffix cgo -o ./bin/main .

# ******************************************************* #

FROM alpine:3.9 AS runtime
# FROM alpine:latest  

# RUN apk add --no-cache rsyslog curl bash sudo ca-certificates
# WORKDIR /root/

COPY --from=build ["/app/bin/main", "/main"] ./
# COPY --from=build /app/static/ /app/static/

ENV GO_ENV=production
EXPOSE ${port}

ENTRYPOINT ["/main"]
