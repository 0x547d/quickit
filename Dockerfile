FROM golang:1.24.5-alpine as builder
ARG app
ARG goProxy
ENV GOPATH=/go
ENV GOOS=linux
ENV CGO_ENABLED=0
ENV GOPROXY=${goProxy}
WORKDIR /go/src/${app}
COPY . .
RUN apk add make git
RUN make binary

FROM alpine:3.22.1 AS runner
LABEL author="Jeyrce Lu<jeyrce@gmail.com>" \
      poweredBy="https://github.com/0x547d/quickit"
RUN echo "Asia/Shanghai" > /etc/timezone && \
    date && \
    mkdir -p /usr/share/zoneinfo/Asia
COPY hack/Shanghai /etc/localtime
COPY hack/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ARG commitId
ARG app
ENV app=${app}
LABEL poweredBy="https://github.com/0x547d/quickit" \
      commitId=${commitId}
WORKDIR /app
COPY hack/ /app
COPY --from=builder /go/src/${app}/build/${app} /usr/local/bin/${app}
EXPOSE 80
VOLUME ["/app/hack/","/etc/timezone:/etc/timezone", "/etc/locatime:/etc/localtime"]
CMD ["sh", "-c", "/usr/local/bin/${app} -c /app/hack/init-config.yml"]
