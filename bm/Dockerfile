FROM golang:1.14-alpine as build

#declare
ENV port ${port:-5000}
WORKDIR /src 

RUN apk add --no-cache git curl

#copy sources
COPY src/* ./ 

#copy dependencies
COPY go.mod ./
# COPY go.sum ./

#install dependencies
RUN go mod download

RUN go build -o bookmaker

# ******************************************************* #

FROM alpine:3.9 as runtime

# WORKDIR /app 

COPY --from=build /src/bookmaker /app/bookmaker
# COPY --from=build /src/static/ /app/static/

EXPOSE ${port}
ENTRYPOINT [ "/app/bookmaker" ]
# CMD [ "sleep" "inf" ]