# build layer
FROM golang:alpine as builder
WORKDIR /build
ADD *.go go.mod ./
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

# end layer
FROM alpine:latest
# copy app
COPY --from=builder /build/app .
# run
CMD [ "./app" ]
