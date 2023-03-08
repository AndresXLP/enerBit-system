FROM golang:1.20-alpine3.17
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download && go build -o main ./cmd
CMD /app/main
