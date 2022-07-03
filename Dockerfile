FROM golang:1.18.3 AS base

WORKDIR /go/src/github.com/MovieStoreGuy/homify

ADD go.mod go.mod
ADD go.sum go.sum

COPY . . 

RUN go mod download

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} make build

FROM gcr.io/distroless/static AS core

COPY --from=base /go/src/github.com/MovieStoreGuy/homify/homify-core /usr/bin/homify-core

CMD [ "/usr/bin/homify-core" ]