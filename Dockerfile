FROM golang as builder
ADD . /go/fibertel_stats/
WORKDIR /go/fibertel_stats
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/fibertel_stats

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/bin/fibertel_stats .
CMD ./fibertel_stats
EXPOSE 9100
