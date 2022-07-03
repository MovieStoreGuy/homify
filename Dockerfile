FROM golang:1.18.3 AS base

WORKDIR /go/src/github.com/MovieStoreGuy/homify

ADD go.mod go.mod
ADD go.sum go.sum

RUN go mod download

COPY . . 

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build --ldflags='-s -w' \ 
        --trimprefix \
        -o homify-core \
        ./cmd/homify

FROM gcr.io/distroless/static AS core

COPY --from=base /go/src/github.com/MovieStoreGuy/homify/homify-core /usr/bin/homify-core

CMD [ "/usr/bin/homify-core" ]