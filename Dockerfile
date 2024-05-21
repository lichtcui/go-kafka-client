FROM debian:bullseye-slim
ADD https://mirrors.aliyun.com/golang/go1.18.8.linux-amd64.tar.gz /tmp/go.tar.gz
RUN tar -C / -xzf /tmp/go.tar.gz; \
    rm /tmp/go.tar.gz;
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
ENV GOPROXY=https://goproxy.cn,direct

RUN apt-get update && apt-get install -y git

WORKDIR /detector
COPY go.* ./
RUN go mod download

COPY ./main.go ./main.go
RUN go build .
