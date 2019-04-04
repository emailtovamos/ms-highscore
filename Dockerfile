# docker build --no-cache -t ms-highscore:latest .
# 1st stage: build Go binary

FROM golang:1.10
WORKDIR /go/src/github.com/emailtovamos/ms-highscore/

# Copy only Go package directories each separately
COPY vendor ./vendor/
COPY ./cli ./cli/
COPY internal ./internal/

ARG CLI_TYPE=server

RUN CGO_ENABLED=0        \
    GOOS=linux           \
    go install           \
      -a                 \
      -installsuffix cgo \
      ./cli/${CLI_TYPE:?}


# 2nd stage: embed Go binary in small Linux distro (== Alpine)

FROM alpine:latest
WORKDIR /app/

# Copy the binary from the first build stage
ARG CLI_TYPE=server
COPY --from=0 /go/bin/${CLI_TYPE} ./binary

CMD ./binary