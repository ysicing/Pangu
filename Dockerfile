# Copyright (c) 2023 ysicing(ysicing.me, ysicing@12306.work) All rights reserved.
# Use of this source code is covered by the following dual licenses:
# (1) Y PUBLIC LICENSE 1.0 (YPL 1.0)
# (2) Affero General Public License 3.0 (AGPL 3.0)
# License that can be found in the LICENSE file.

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
