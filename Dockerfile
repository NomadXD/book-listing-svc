FROM golang:1.19.0-alpine as build-env

RUN mkdir /app
WORKDIR /app
COPY go.mod ./

# Create a new user with UID 10014
RUN addgroup -g 10014 choreo && \
    adduser  --disabled-password  --no-create-home --uid 10014 --ingroup choreo choreouser


# COPY the source code as the last step
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app -buildvcs=false

FROM alpine
COPY --from=build-env /go/bin/app /go/bin/app
COPY books.json /tmp/data/books.json

EXPOSE 8080

USER 10014
ENTRYPOINT ["/go/bin/app"]