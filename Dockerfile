FROM golang:1.15
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 5678
ENTRYPOINT ["/go/bin/cmpt433-frontend"]
