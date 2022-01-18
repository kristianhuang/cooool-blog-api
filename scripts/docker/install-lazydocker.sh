#!/bin/bash

#
# Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.
#

sudo wget https://github.com/jesseduffield/lazydocker/releases/download/v0.12/lazydocker_0.12_Linux_x86_64.tar.gz

sudo mkdir ./lazydocker \
     && mv ./lazydocker*.tar.gz ./lazydocker \
     && tar -zxvf ./lazydocker/lazydocker*.tar.gz  ./lazydocker