FROM ysicing/god AS builder

WORKDIR /app

RUN go install github.com/mitchellh/gox@latest \
    && go install github.com/go-task/task/v3/cmd/task@latest

COPY go.mod ./

COPY go.sum ./

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod download

COPY . ./

RUN task

FROM ysicing/debian

WORKDIR /app

COPY --from=builder /app/_output/pangu_linux_amd64 /app/pangu

CMD ["/app/pangu", "server"]
