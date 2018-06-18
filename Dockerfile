FROM golang:1.10-alpine3.7 as builder

RUN mkdir -p /go/src/github.com/dsalazar32/go-certbotify-aws
COPY . /go/src/github.com/dsalazar32/go-certbotify-aws
RUN cd /go/src/github.com/dsalazar32/go-certbotify-aws && go install .

FROM certbot/dns-route53
COPY --from=builder /go/bin/go-certbotify-aws /bin/go-certbotify-aws
ENTRYPOINT ["go-certbotify-aws"]
