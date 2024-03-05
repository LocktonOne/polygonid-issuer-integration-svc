FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokene/polygonid-issuer-integration
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/polygonid-integration-svc /go/src/gitlab.com/tokene/polygonid-issuer-integration


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/polygonid-integration-svc /usr/local/bin/polygonid-integration-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["polygonid-integration-svc"]
