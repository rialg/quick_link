#! /usr/bin/env bash

# Build docker image
docker image build . --tag quick_link
# Run the quicklink proxy
docker run --mount type=bind,source=/usr/local/redirects.json,target=$(pwd)/static/redirects.json \
           -d --restart="always" --name "quick_link" quick_link:latest
# Get the container IP addr
IP_ADDR=$(docker container inspect quick_link -f '{{ .NetworkSettings.IPAddress }}')

# Add the quick link hostname
echo "${IP_ADDR}	<QUICKLINK_ROOT>" >> /etc/hosts
