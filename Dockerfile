FROM golang:alpine AS Go

RUN mkdir /go/main
ADD main.go /go/main/
RUN cd /go/main && go mod init main && go mod tidy && go build

FROM scratch

COPY --from=Go /go/main/main /main
CMD ["/main"]
