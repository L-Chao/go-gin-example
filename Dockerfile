FROM scratch

WORKDIR $GOPATH/src/go-gin-example
COPY . $GOPATH/src/go-gin-example

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]