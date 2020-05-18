FROM golang

RUN go get github.com/erif93/TAPERA-Mobile-BackEnd-Preparation

ADD . /go/src/github.com/erif93/TAPERA-Mobile-BackEnd-Preparation

RUN go install github.com/erif93/TAPERA-Mobile-BackEnd-Preparation

ENTRYPOINT /go/bin/TAPERA-Mobile-BackEnd-Preparation

EXPOSE 8080