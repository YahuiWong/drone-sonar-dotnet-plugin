FROM golang:1.13.4-alpine as build
RUN mkdir -p /go/src/github.com/yahuiwong/drone-sonar-dotnet-plugin
WORKDIR /go/src/github.com/yahuiwong/drone-sonar-dotnet-plugin 
COPY *.go ./
COPY vendor ./vendor/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sonar

FROM mcr.microsoft.com/dotnet/sdk:6.0-alpine3.14-amd64
# RUN sed -i "s/http:\/\/deb.debian.org\/debian/https:\/\/mirrors.aliyun.com\/debian/g" /etc/apt/sources.list
# RUN apt clean \
#     && apt update \
#     && apt install default-jdk -y \
#     && apt clean
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update \
&& apk add openjdk8 nodejs \
&& rm -rf /tmp/* /var/cache/apk/*

COPY --from=build /go/src/github.com/yahuiwong/drone-sonar-dotnet-plugin/drone-sonar /bin/
WORKDIR /bin

RUN dotnet tool install --global dotnet-sonarscanner

ENV PATH $PATH:/root/.dotnet/tools

ENTRYPOINT /bin/drone-sonar
