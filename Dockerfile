# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang
MAINTAINER Dmitry Zeldin <dmitry@zeldin.pro>

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/ZelJin/bugurtobot
# Install dependencies via glide
RUN curl https://glide.sh/get | sh
RUN cd /go/src/github.com/ZelJin/bugurtobot && glide install

# Build the app inside the container.
RUN go install github.com/ZelJin/bugurtobot

# Run the app by default when the container starts.
ENTRYPOINT /go/bin/bugurtobot
