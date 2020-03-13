ARG FROMBUILD=golang:alpine
ARG FROM=alpine
FROM $FROMBUILD as build
ARG GOARCH
RUN apk add --no-cache git
RUN go get -d github.com/subspacecommunity/subspace && \
    cd /go/src/github.com/subspacecommunity/subspace && \
    go generate && \
    CGO_ENABLED=0 GOARCH=$GOARCH go build -v --compiler gc --ldflags "-extldflags -static -s -w" -o /go/bin/subspace

FROM $FROM
ARG GOARCH
LABEL maintainer="squat <lserven@gmail.com>"
RUN echo -e "https://dl-3.alpinelinux.org/alpine/edge/main\nhttps://dl-3.alpinelinux.org/alpine/edge/community" > /etc/apk/repositories && \
    apk add --no-cache wireguard-tools perl
ENV KILOSUBSPACE_WG /usr/bin/wg
COPY bin/$GOARCH/kilosubspace /usr/local/bin/kilosubspace
RUN ln -s /usr/local/bin/kilosubspace /usr/local/bin/wg
COPY --from=build /go/bin/subspace /usr/local/bin/subspace
ENTRYPOINT ["/usr/local/bin/subspace"]
