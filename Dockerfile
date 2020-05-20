FROM golang:latest 

RUN go get github.com/gin-gonic/gin

RUN mkdir /app 

ADD . /app/ 

WORKDIR /app 

RUN go build -o main . 

CMD ["/app/main"]


# FROM golang:1.12-alpine

# RUN apk add --no-cache git

# RUN go get github.com/erif93/TAPERA-Mobile-BackEnd-Preparation

# ADD . /go/src/github.com/erif93/TAPERA-Mobile-BackEnd-Preparation

# RUN go install github.com/erif93/TAPERA-Mobile-BackEnd-Preparation

# ENTRYPOINT /go/bin/TAPERA-Mobile-BackEnd-Preparation

# EXPOSE 8080

# FROM golang:1.12-alpine
# RUN apk add --no-cache git
# # Set the Current Working Directory inside the container
# WORKDIR /app/tapera-sim
# # We want to populate the module cache based on the go.{mod,sum} files.
# # COPY go.mod .
# # COPY go.sum .
# RUN go mod download
# COPY . .
# # Build the Go app
# RUN go build -o ./out/tapera-sim .

# # This container exposes port 8080 to the outside world
# EXPOSE 8080
# # Run the binary program produced by `go install`
# CMD ["./out/tapera-sim"]