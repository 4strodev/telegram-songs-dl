## This image uses multistagging
## This step just builds the bot
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS build

ARG TARGETOS
ARG TARGETARCH
WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /go/bin/app /go/src/app/cmd

## Using distroless images to reduce image size
FROM python:3.12-slim

RUN apt-get update && apt-get install -y bash curl ffmpeg && pip install yt-dlp

COPY --from=build /go/bin/app /
CMD ["/app"]
