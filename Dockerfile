FROM golang:1.13.4-alpine as build
RUN mkdir -p /go/src/github.com/aosapps/drone-sonar-plugin
WORKDIR /go/src/github.com/aosapps/drone-sonar-plugin 
COPY *.go ./
COPY vendor ./vendor/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-sonar

FROM mcr.microsoft.com/dotnet/sdk:3.1



RUN apt-get update \
    && apt-get install default-jdk -y \
    && apt-get clean

COPY --from=build /go/src/github.com/aosapps/drone-sonar-plugin/drone-sonar /bin/
WORKDIR /bin

RUN dotnet tool install --global dotnet-sonarscanner

ENV PATH $PATH:/root/.dotnet/tools

ENTRYPOINT /bin/drone-sonar
