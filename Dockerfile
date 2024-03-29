FROM golang:1.21-alpine as build
WORKDIR /go/src/app
COPY . /go/src/app
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10
LABEL maintainer="Luke Tainton <luke@tainton.uk>"
LABEL org.opencontainers.image.source="https://github.com/luketainton/iPilot"
COPY --from=build /go/bin/app /
CMD ["/app"]
