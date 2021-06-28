FROM golang:latest as builder
WORKDIR /app
ADD . /app/
#RUN go mod tidy
RUN CGO_ENABLED=0 go build -o booklibrary

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache bash
COPY --from=builder /app/booklibrary .
COPY --from=builder /app/static/* static/
COPY --from=builder /app/templates/* templates/

CMD ["./booklibrary"]