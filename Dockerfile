FROM golang:1.12.5-alpine
RUN apk add git

ENV GO111MODULE=on

RUN mkdir /assistant
WORKDIR /assistant

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# RUN go get -d -u -v ./...
RUN go install -v ./...


CMD ["assistant"]
