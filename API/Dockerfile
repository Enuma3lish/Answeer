
FROM golang:1.22.3 as builder

WORKDIR /app

# copy go.mod 
COPY go.mod ./

#  download dependency
RUN go mod download

COPY . .


RUN go mod tidy
RUN go build -o main .

FROM alpine:latest

RUN apk add --no-cache tzdata
ENV TZ=Asia/Taipei

# setting world folder
WORKDIR /root/

# 从构建阶段复制造建好的二进制文件
COPY --from=builder /app/main .

# Ensure can execute
RUN chmod +x ./main

# Setting port
EXPOSE 8080

# start
CMD ["./main"]
