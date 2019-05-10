FROM golang
ADD . app
WORKDIR app
RUN go build -mod=vendor
EXPOSE 8090
ENTRYPOINT ["./taas"]