#!/usr/bin/env bash

#
# Copyright 2021 Kristian Huang <kristianhuang@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.
#

CONFIG_PATH=$(dirname "${BASH_SOURCE[0]}")/../../config
CONFIG=$CONFIG_PATH/blog-apiserver.yaml
CONFIG_LOCAL=$CONFIG_PATH/local-blog-apiserver.yaml

# If you need new config.
# cp $CONFIG $CONFIG_LOCAL


gowatch -p ./cmd/apiserver -args="-c=$CONFIG_LOCAL"