FROM golang:alpine as build-backend

ARG VERSION

COPY . /srv/backend
RUN cd /srv/backend && \
    go build -o app -ldflags "-X main.revision=${VERSION} -s -w" .

FROM alpine

COPY --from=build-backend /srv/backend/app /srv/app
RUN mkdir -p /home/app && \
    adduser -s /bin/sh -D -u 1001 app && chown -R app:app /home/app
RUN chown -R app:app /srv && \
    chmod +x /srv/app

WORKDIR /srv

CMD ["/srv/app"]
