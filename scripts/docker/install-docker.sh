#!/bin/bash

#
# Copyright 2021 KristianHuang <kristianhuang007@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.
#

# remove old versions
sudo yum remove docker \
  docker-client \
  docker-client-latest \
  docker-common \
  docker-latest \
  docker-latest-logrotate \
  docker-logrotate \
  docker-engine

sudo yum install -y yum-utils

# add yum repo
sudo yum-config-manager \
  --add-repo \
  https://download.docker.com/linux/centos/docker-ce.repo

# install docker
sudo yum install -y docker-ce docker-ce-cli containerd.io

# add ali mirrors
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://zjwcef2k.mirror.aliyuncs.com"]
}
EOF

# if you want to edit the default data path
# {
#   "registry-mirrors": ["https://zjwcef2k.mirror.aliyuncs.com"],
#   "graph": "path" # centos
#   "data-root": "path" # ubuntu
# }

# install docker-compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# reload docker
sudo systemctl daemon-reload
sudo systemctl restart docker