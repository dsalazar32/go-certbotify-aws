FROM golang:1.10-alpine3.7 as builder

RUN mkdir -p /go/src/github.com/dsalazar32/go-gen-ssl
COPY . /go/src/github.com/dsalazar32/go-gen-ssl
RUN cd /go/src/github.com/dsalazar32/go-gen-ssl && go install .

FROM certbot/dns-route53
COPY --from=builder /go/bin/go-gen-ssl /bin/go-gen-ssl
ENTRYPOINT ["go-gen-ssl"]
