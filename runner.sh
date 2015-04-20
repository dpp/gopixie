#!/bin/bash

sudo docker build -t gopixie builder

sudo docker run --rm -t -i -v $(pwd):/data gopixie /data/scripts/start
