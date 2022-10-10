FROM ubuntu:20.04

SHELL ["/bin/bash", "-c"]

RUN apt update && apt install -y default-jdk && apt install -y curl && apt -y install zip unzip

RUN curl -s https://get.sdkman.io | bash

RUN source /root/.sdkman/bin/sdkman-init.sh && \
    sdk install kotlin