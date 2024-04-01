# Build stage
FROM docker.io/golang:1.22
WORKDIR /src
COPY go.mod go.sum /src/
RUN go mod download && go mod verify

COPY . /src/
RUN go build -o /go-api-template

EXPOSE 8080

CMD [ "/go-api-template" ]
