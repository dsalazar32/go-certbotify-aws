FROM golang:1.10 as builder

RUN mkdir -p /go/src/github.com/dsalazar32/go-certbotify-aws
COPY . /go/src/github.com/dsalazar32/go-certbotify-aws
RUN cd /go/src/github.com/dsalazar32/go-certbotify-aws && CGO_ENABLED=0 go install .

FROM certbot/dns-route53
COPY --from=builder /go/bin/go-certbotify-aws /bin/go-certbotify-aws
ENTRYPOINT ["go-certbotify-aws"]
