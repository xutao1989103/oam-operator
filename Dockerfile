FROM alpine:3.11

RUN addgroup -S oam \
    && adduser -S -g oam oam \
    && apk --no-cache add ca-certificates

WORKDIR /home/oam

COPY /bin/oamOperator .

RUN chown -R oam:oam ./

USER oam

ENTRYPOINT ["./oamOperator"]