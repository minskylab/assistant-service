FROM golang:1.14

RUN mkdir /assistant
WORKDIR /assistant

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install -v ./...

CMD ["ctl"]
