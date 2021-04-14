FROM alpine:latest

RUN apk update && apk add --no-cache libc6-compat
COPY ./jobs ./go/src/

CMD ["/go/src/jobs"]