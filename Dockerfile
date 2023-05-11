FROM golang:1.20-bullseye as build
# FROM golang:1.14-alpine as build
# FROM cosmtrek/air

#declare
ENV port ${port:-5000}
WORKDIR /app 

RUN apt install git curl

#copy dependencies
COPY go.mod /app 
COPY go.sum /app 

#install dependencies
RUN go mod download
# RUN go get github.com/githubnemo/CompileDaemon

COPY ./app/* .

# Unit tests
# RUN CGO_ENABLED=0 go test -v

# ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main

# RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \  
#     && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN go build -o ./out/bookmaker .

# ******************************************************* #

FROM alpine:3.9 as runtime
# WORKDIR /app 
RUN apk add ca-certificates
COPY --from=build /app/out/bookmaker /app/bookmaker
# COPY --from=build /app/static/ /app/static/

EXPOSE ${port}
# ENTRYPOINT [ "/app/bookmaker" ]
# CMD [ "sleep" "inf" ]