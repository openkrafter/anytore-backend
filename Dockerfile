# For build
FROM golang:1.21.5

## update packages
RUN apt-get update
RUN apt-get upgrade -y

## install packages
### -

## build application
RUN mkdir -p /go/src/anytore-backend
COPY ./ /go/src/anytore-backend/

WORKDIR /go/src/anytore-backend
RUN go build cmd/anytore-backend/main.go


# For release
FROM golang:1.21.5

## update packages
RUN apt-get update
RUN apt-get upgrade -y

## release application
COPY --from=0 /go/src/anytore-backend/main /root/

## environment variable
### anytore
#### LOG_LEVEL: info or debug
ENV LOG_LEVEL info

### Gin
#### GIN_MODE: release or debug
ENV GIN_MODE release

ENV PORT 80

### AWS
#### AWS_REGION: ap-northeast-1 or us-east-1 or ...
ENV AWS_REGION ap-northeast-1

## Dockerfile common
EXPOSE 80

CMD /root/main
