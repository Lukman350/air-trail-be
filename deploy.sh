#!/usr/bin/env bash
set -e

if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
  echo "[ENV] .env file loaded"
else
  echo "[ENV] ⚠️ No .env file found, using defaults"
fi

APP_NAME="air_trail_backend"
DOCKER_IMAGE="$APP_NAME"
DOCKER_TAG=${DOCKER_TAG:-latest}
APP_PORT=${APP_PORT:-8080}

echo "[DOCKER] Building Docker image: $DOCKER_IMAGE:$DOCKER_TAG"
docker build -t $DOCKER_IMAGE:$DOCKER_TAG .
echo "[DOCKER] ✅ Docker image created: $DOCKER_IMAGE:$DOCKER_TAG"

clear

echo "[DOCKER] Removing old container if exists..."
docker rm -f $APP_NAME 2>/dev/null || true

echo "[DOCKER] Starting the application..."
docker run \
	--add-host=host.docker.internal:host-gateway \
	--log-driver=json-file \
	--log-opt max-size=10m \
	--log-opt max-file=3 \
	-d --name $APP_NAME -p $APP_PORT:$APP_PORT $DOCKER_IMAGE:$DOCKER_TAG
echo "[DOCKER] ✅ Application started on port $APP_PORT"
