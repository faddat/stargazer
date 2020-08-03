FROM alpine:3.12
RUN apk add -U --no-cache ca-certificates


COPY ./build/stakewatcher /usr/bin/stakewatcher
EXPOSE 8080

CMD ["stakewatcher"]
