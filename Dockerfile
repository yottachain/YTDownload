FROM registry.cn-beijing.aliyuncs.com/ytc-common/alpine:3
LABEL maintainer="yuanye@yottachain.io"
LABEL desc="download service"
LABEL src="https://github.com/yottachain/YTDownload.git"

WORKDIR /app
COPY ./download /app/download

ENV GIN_MODE=release

EXPOSE 8081

ENTRYPOINT ["/app/download"]
