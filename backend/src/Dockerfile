FROM golang
ADD . app
WORKDIR app
RUN go build -mod=vendor
EXPOSE 9090
ENTRYPOINT ["./taas"]