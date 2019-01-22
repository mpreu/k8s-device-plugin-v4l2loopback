## Build stage
FROM golang:1.10.3 as build

WORKDIR /go/src/github.com/mpreu/k8s-device-plugin-v4l2loopback
COPY . .
# Install dependencies if necessary
#RUN go get github.com/golang/dep/cmd/dep
#COPY Gopkg.toml Gopkg.lock ./
#RUN dep ensure -v -vendor-only
# Build and install
RUN go install -v github.com/mpreu/k8s-device-plugin-v4l2loopback

## Run stage
FROM golang:1.10.3-alpine3.8
COPY --from=build /go/bin/k8s-device-plugin-v4l2loopback /bin/k8s-device-plugin-v4l2loopback

CMD ["/bin/k8s-device-plugin-v4l2loopback"]