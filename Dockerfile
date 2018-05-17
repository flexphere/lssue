FROM library/node:8.11.1-alpine AS build-ui

ARG env=development
ENV NODE_ENV=$env

WORKDIR /go/src/github.com/flexphere/lssue/static

COPY ./static .

RUN set -e \
	&& rm -rf node_modules dist \
	&& mkdir dist \
    && yarn install \
	&& yarn build

#------------------------------------------------------------

FROM golang:latest AS build-app

WORKDIR /go/src/github.com/flexphere/lssue

RUN set -e \
	&& rm -rf vendor \
	&& go get -u github.com/golang/dep/cmd/dep

COPY . .

RUN set -e \
	&& dep ensure \
	&& go build . 

# ------------------------------------------------------------

FROM debian:latest

WORKDIR /lssue

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN mkdir -p /lssue/static/dist
COPY --from=build-app /go/src/github.com/flexphere/lssue/lssue lssue
COPY --from=build-ui /go/src/github.com/flexphere/lssue/static/dist static/dist

CMD ["/lssue/lssue"]