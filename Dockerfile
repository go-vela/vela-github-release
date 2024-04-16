# SPDX-License-Identifier: Apache-2.0

# set a global Docker argument for the default CLI version
#
# https://github.com/moby/moby/issues/37345
ARG GH_VERSION=2.14.4

###################################################################################
##    docker build --no-cache --target binary -t vela-github-release:binary .    ##
###################################################################################

FROM alpine:3.19.1@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b as binary

ARG GH_VERSION

ADD https://github.com/cli/cli/releases/download/v${GH_VERSION}/gh_${GH_VERSION}_linux_amd64.tar.gz /tmp/gh.tar.gz

RUN tar -xzf /tmp/gh.tar.gz -C /bin

RUN cp /bin/gh_${GH_VERSION}_linux_amd64/bin/gh /bin/gh

RUN chmod 0700 /bin/gh

##################################################################
##    docker build --no-cache -t vela-github-release:local .    ##
##################################################################

FROM alpine:3.19.1@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b

ARG GH_VERSION

ENV PLUGIN_GH_VERSION=${GH_VERSION}

RUN apk add --update --no-cache git ca-certificates

COPY --from=binary /bin/gh /bin/gh

COPY release/vela-github-release /bin/vela-github-release

ENTRYPOINT [ "/bin/vela-github-release" ]
