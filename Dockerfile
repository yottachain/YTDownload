FROM harbor1-c3-bj.yottachain.net/yt-common/nginx:alpine
LABEL maintainer="yuanye@yottachain.io"
LABEL desc="download service"
LABEL src="https://github.com/yottachain/YTDownload.git"

WORKDIR /app
COPY ./DownloadNew /app/DownloadNew

ENV GIN_MODE=release

EXPOSE 8081

ENTRYPOINT ["/app/DownloadNew"]