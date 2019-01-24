## Build stage
FROM golang:1.10.3 as build

# Install go dep
RUN go get -u github.com/golang/dep/...

WORKDIR /go/src/github.com/mpreu/k8s-device-plugin-v4l2loopback

COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./

# Install dependencies if necessary
#RUN go get github.com/golang/dep/cmd/dep
#COPY Gopkg.toml Gopkg.lock ./
#RUN dep ensure -v -vendor-only
# Build and install
RUN go install -v github.com/mpreu/k8s-device-plugin-v4l2loopback

## Run stage
FROM golang:1.10.3
COPY --from=build /go/bin/k8s-device-plugin-v4l2loopback /bin/k8s-device-plugin-v4l2loopback

ENTRYPOINT ["/bin/k8s-device-plugin-v4l2loopback"]