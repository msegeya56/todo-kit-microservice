# First stage: build stage
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /go/src/github.com/msegeya56/todo-kit-microsrvice/customer



ADD . /go/src/github.com/msegeya56/todo-kit-microsrvice/customer/

# Copy go mod and sum files
COPY go.mod . 
COPY go.sum .

# Download and install Go module dependencies
RUN go mod download

# Copy the entire Go application directory to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/github.com/msegeya56/todo-kit-microservice/customer/cmd/main.go github.com/msegeya56/todo-kit-microservice/customer/cmd/maain.go

# Second stage: final stage
FROM alpine:latest

# Set the working directory inside the container (a generic path for multiple microservices)
WORKDIR /ecommerce_platform

# Copy the built binary from the builder stage to the final stage
COPY --from=builder /go/bin/todo-kit-microservice/customer/cmd/main.go /todo-kit-microservice/

# Set the entrypoint to run the binary
ENTRYPOINT ["/todo-kit-microsrvice/custome/cmd/main.gor"]



















RUN mkdir -p /go/src/github.com/msegeya56/todo-kit-microsrvice/customer




# RUN go get  -t -v ./...
# RUN go get  github.com/canthefason/go-watcher








# ENTRYPOINT  watcher -run github.com/msegeya56/todo-kit-microsrvice/customer/cmd  -watch github.com/msegeya56/todo-kit-microsrvice/customer






# RUN go get  -t -v ./...
# RUN go get  github.com/canthefason/go-watcher
# RUN go install github.com/canthefason/go-watcher/cmd/watcher
# WORKDIR /go/src/github.com/msegeya56/artery_platform/arts

# RUN go mod vendor

# RUN go build github.com/msegeya56/artery_platform/arts/cmd/main.go 

#ENTRYPOINT  [watcher -run github.com/msegeya56/artery_platform/arts/cmd  -watch github.com/msegeya56/artery_platform/arts]







