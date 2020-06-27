# build stage
FROM golang:1.14-alpine as build

WORKDIR $GOPATH/app/

RUN apk add git

# copy and download dependencies
COPY go.* ./
RUN go mod download

#compile app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main

#resulting app
FROM scratch as final
COPY --from=build go/app/main /app/
WORKDIR /app
EXPOSE 8080
ENTRYPOINT [ "./main" ]