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
FROM gcr.io/distroless/static-debian12:latest

COPY --from=build /go/bin/app /
CMD ["/app"]
