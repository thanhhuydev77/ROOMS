FROM golang:1.14

RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o main .
CMD ["/app/main"]