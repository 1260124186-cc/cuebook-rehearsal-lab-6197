#!/bin/bash
# Build the delivery image locally. The optional second argument selects the target platform.
set -euo pipefail

IMAGE_NAME=${1:-my-go-task}
PLATFORM=${2:-linux/amd64}

docker build --platform "$PLATFORM" -f benzhi.Dockerfile -t "$IMAGE_NAME" .

echo ""
echo "✅ Docker image '$IMAGE_NAME' built successfully!"
echo "▶ Start service: docker run --rm -d -p 8080:8080 $IMAGE_NAME"
echo "♥ Check health:  curl -fsS http://127.0.0.1:8080/healthz"
