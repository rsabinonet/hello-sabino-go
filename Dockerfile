FROM alpine:latest

MAINTAINER Rodrigo Sabino (rsabino@gmail.com)

WORKDIR "/opt"

ADD .docker_build/hello-sabino /opt/bin/hello-sabino
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/hello-sabino"]

