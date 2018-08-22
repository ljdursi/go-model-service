# Build stage
FROM golang

ENV DIR=/go/src/github.com/CanDIG/go-model-service
WORKDIR ${DIR}
ADD . ${DIR}

# Install CLI tool dependencies for install.sh script
RUN ./install_dep.sh
RUN ./install.sh

EXPOSE 3000

ENTRYPOINT "${DIR}/main" --port=3000
