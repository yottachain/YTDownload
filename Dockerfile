FROM harbor1-c3-bj.yottachain.net/yt-common/alpine:3
LABEL maintainer="yuanye@yottachain.io"
LABEL desc="download service"
LABEL src="https://github.com/yottachain/YTDownload.git"

WORKDIR /app
COPY ./DownloadNew /app/DownloadNew

ENV GIN_MODE=release

EXPOSE 80810

ENTRYPOINT ["/app/DownloadNew"]