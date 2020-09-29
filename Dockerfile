FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir /app

RUN apk add --no-cache \
        libc6-compat

WORKDIR /app

COPY sf1.exe /app

COPY view/public /app/view/public

CMD ["./sf1.exe"]

EXPOSE 9090
