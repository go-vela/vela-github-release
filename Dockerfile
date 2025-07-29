# SPDX-License-Identifier: Apache-2.0

# set a global Docker argument for the default CLI version
#
# https://github.com/moby/moby/issues/37345
# renovate: datasource=github-tags depName=cli/cli extractVersion=^v(?<version>.*)$
ARG GH_VERSION=2.76.1

###################################################################################
##    docker build --no-cache --target binary -t vela-github-release:binary .    ##
###################################################################################

FROM alpine:3.22.1@sha256:4bcff63911fcb4448bd4fdacec207030997caf25e9bea4045fa6c8c44de311d1 as binary

ARG GH_VERSION

ENV GH_RELEASE_URL="https://github.com/cli/cli/releases/download/v${GH_VERSION}"
ENV GH_FILENAME="gh_${GH_VERSION}_linux_amd64.tar.gz"
ENV GH_CHECKSUM_FILENAME="gh_${GH_VERSION}_checksums.txt"

RUN wget -q "${GH_RELEASE_URL}/${GH_FILENAME}" -O "${GH_FILENAME}" && \
  wget -q "${GH_RELEASE_URL}/${GH_CHECKSUM_FILENAME}" -O "${GH_CHECKSUM_FILENAME}" && \
  grep "${GH_FILENAME}" "${GH_CHECKSUM_FILENAME}" | sha256sum -c && \
  tar -xf "${GH_FILENAME}" && \
  mv "${GH_FILENAME%.tar.gz}/bin/gh" /bin/gh && \
  chmod 0700 /bin/gh


##################################################################
##    docker build --no-cache -t vela-github-release:local .    ##
##################################################################

FROM alpine:3.22.1@sha256:4bcff63911fcb4448bd4fdacec207030997caf25e9bea4045fa6c8c44de311d1

ARG GH_VERSION

ENV PLUGIN_GH_VERSION=${GH_VERSION}

RUN apk add --update --no-cache git ca-certificates

COPY --from=binary /bin/gh /bin/gh

COPY release/vela-github-release /bin/vela-github-release

ENTRYPOINT [ "/bin/vela-github-release" ]
