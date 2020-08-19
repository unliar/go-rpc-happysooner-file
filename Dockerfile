FROM golang:latest as go-build

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN GO111MODULE=on GOPROXY=https://goproxy.io CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/app.go


FROM orvice/go-runtime:lite

WORKDIR /

ENV MICRO_REGISTRY=consul

ENV MICRO_REGISTRY_ADDRESS=consul:8500

COPY --from=go-build /app/main /

ENTRYPOINT ["/main"]