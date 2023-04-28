FROM registry.cn-beijing.aliyuncs.com/haloop-base/go-build:1.19.5 as builder
ADD . /go/src/app
WORKDIR /go/src/app
RUN go mod tidy && go build -ldflags="-s -w" -o api-server .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/api-server .
EXPOSE 8080
CMD ["./api-server"]



