GIT_SHA=$(git rev-parse --short HEAD) \
GIT_LABEL=$(git rev-parse --abbrev-ref HEAD) \
docker-compose build && \
docker-compose push
