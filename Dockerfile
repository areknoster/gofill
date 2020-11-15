FROM golang:latest
WORKDIR /go/src/github.com/areknoster/gofill
RUN go get fyne.io/fyne
COPY  ./   ./
RUN CGO_ENABLED=0 GOOS=darwin go build -o app GoFill

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/alexellis/href-counter/app .
CMD ["./app"]