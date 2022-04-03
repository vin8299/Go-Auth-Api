    FROM golang:1.16-alpine  
    RUN mkdir /go-auth
    ADD . /go-auth
    WORKDIR /go-auth
    COPY go.mod .
    COPY go.sum .
    RUN go mod download
    COPY . .

    RUN go build -o authmain .
    CMD ["/go-auth/authmain"]