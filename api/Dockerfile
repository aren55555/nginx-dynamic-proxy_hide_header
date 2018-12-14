FROM golang:1.11-stretch

# Install OS libs
RUN apt-get update -y && \
  apt-get install -y \
    curl \
    git \
    wget \
    xmlsec1
# Install Go's dependency tool
RUN go get -u github.com/golang/dep/cmd/dep

###################
# App Preparation #
###################
# Set go working dir for this app
WORKDIR /go/src/github.com/aren55555/nginx-dynamic-proxy_hide_header/api
# Copy dep files
COPY Gopkg.lock .
COPY Gopkg.toml .

########################
# SSH Config for Stash #
########################
# Get dependencies
RUN dep ensure \
    -v \
    -vendor-only

###################
# App Preparation #
###################
# Copy all of the app
COPY . .
# Compile binary, and put in PATH
RUN go install -v ./...
# Exec go bin
EXPOSE 8080
CMD ["api"]
