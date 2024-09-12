FROM golang:alpine

WORKDIR /app

COPY . .

RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go build -o main main.go

CMD [ "./main" ]