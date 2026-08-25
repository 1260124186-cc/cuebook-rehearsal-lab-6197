# Delivery image: compile the CLI once, then start its HTTP health interface by default.
FROM golang:1.26.2

WORKDIR /app
COPY go.mod ./
COPY . ./
RUN go build -o /usr/local/bin/cuebook ./cmd/cuebook

ENV PORT=8080
EXPOSE 8080
CMD ["cuebook", "serve"]

# Single-platform local build example:
# docker build --platform linux/amd64 -f benzhi.Dockerfile -t <image> .
