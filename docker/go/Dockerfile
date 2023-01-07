FROM golang

SHELL ["/bin/bash", "-c"]

RUN curl -fsSL \
        https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
        sh

RUN echo 'export GO_PATH=/go' >> ~/.bash_profile \
    && echo 'export PATH=$PATH:$GO_PATH/bin' >> ~/.bash_profile \
    && source ~/.bash_profile \
