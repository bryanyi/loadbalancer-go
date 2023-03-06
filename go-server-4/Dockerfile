FROM golang:1.19.3

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

RUN go build -o /main main.go

EXPOSE 5004

ENTRYPOINT [ "/main" ]
