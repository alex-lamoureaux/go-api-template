# Build stage
FROM golang:1.22 as build
WORKDIR /src
COPY go.mod go.sum /src/
RUN go mod download && go mod verify

COPY . /src/
RUN go build -o /go-api-template

# Run stage
FROM golang:1.22-alpine as run
COPY --from=build /go-api-template /go-api-template

EXPOSE 8080

CMD [ "/go-api-template" ]
