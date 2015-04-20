# A Dockerfile for compiling Pixie *and* Compiling Go
# And running both...

FROM       ubuntu:14.04
MAINTAINER David Pollak <feeder.of.the.bears@gmail.com>

RUN apt-get update && apt-get --yes install curl pypy python-pip libffi-dev libedit-dev libuv-dev libboost-all-dev build-essential git bison gcc g++ libboost-dev pkg-config make bzip2 python2.7

RUN echo 'bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)' | /bin/bash

RUN echo 'source /root/.gvm/scripts/gvm && gvm install go1.4 && gvm use go1.4' | /bin/bash

RUN pip install cffi
