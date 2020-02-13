FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
WORKDIR /go/src/github.com/yoshihisamurakami/hr.editorial
COPY . .
WORKDIR /go/src/github.com/yoshihisamurakami/hr.editorial/web
RUN go build -o ../bin/web
WORKDIR /go/src/github.com/yoshihisamurakami/hr.editorial/crawler
RUN go build -o ../bin/crawler

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/yoshihisamurakami/hr.editorial /app
COPY --from=builder /go/src/github.com/yoshihisamurakami/hr.editorial/templates/ .

CMD /app/bin/web 
