FROM alpine:3.18.4

ARG GIT_HASH
ARG GIT_BRANCH
LABEL gitHash=$GIT_HASH gitBranch=$GIT_BRANCH

WORKDIR /
COPY cmd/zeus-server /zeus-server

CMD ["/zeus-server"]
