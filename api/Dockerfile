# syntax=docker/dockerfile:1

FROM golang:1.22.6

# Set destination for COPY
WORKDIR /app

# Def. ENV Vars
ENV GIN_MODE=release

# Download Go modules
COPY . .

RUN cd album && go mod download && go mod tidy
RUN cd client && go mod download && go mod tidy
RUN cd transaction && go mod download && go mod tidy
RUN go mod download && go mod tidy


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /api

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["/api"]