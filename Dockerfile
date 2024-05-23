FROM golang:1.22-alpine AS builder
# https://github.com/chaseSpace/k8s-tutorial-cn
# 缓存依赖
WORKDIR /go/cache
COPY go.mod .
# COPY go.sum .
RUN GOPROXY=https://goproxy.cn,direct go mod tidy

WORKDIR /build
COPY . .

# 关闭cgo的原因：使用了多阶段构建，go程序的编译环境和运行环境不同，不关就无法运行go程序
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GO111MODULE=auto go build -ldflags "-w -extldflags -static" -o main

#FROM scratch as prod
FROM alpine as prod
# 通过 http://www.asznl.com/post/48 了解docker基础镜像：scratc、busybox、alpine
# 比他们还小的是distroless   由谷歌提供，了解：https://github.iotroom.top/GoogleContainerTools/distroless

# alpine设置时区
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories &&  \
    apk add -U tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata && date

COPY --from=builder /build/main .

EXPOSE 3000
ENTRYPOINT ["/main"]