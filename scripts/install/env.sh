#!/usr/bin/env bash

#
# Copyright 2021 Kristian Huang <kristianhuang007@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.
#

# 源码根目录
ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

# 设置统一的密码，方便记忆
readonly PASSWORD=${PASSWORD:-'Y%zsrS4b'}

# 生成文件存放目录
LOCAL_OUTPUT_ROOT="${ROOT}/${OUT_DIR:-_output}"

# 安装目录
readonly INSTALL_DIR=${INSTALL_DIR:-/tmp/installation}
mkdir -p ${INSTALL_DIR}
readonly ENV_FILE=${ROOT}/scripts/install/env.sh

# Mysql 配置信息
readonly MYSQL_ADMIN_USERNAME=${MYSQL_ADMIN_USERNAME:-root} # MYSQL root 用户
readonly MYSQL_ADMIN_PASSWORD=${MYSQL_ADMIN_PASSWORD:-${PASSWORD}} # MYSQL root 用户密码
readonly MYSQL_HOST=${MYSQL_HOST:-127.0.0.1:3306} # MYSQL 主机地址
readonly MYSQL_DATABASE=${MYSQL_DATABASE:-blog} # MYSQL blog 应用使用的数据库名
readonly MYSQL_USERNAME=${MYSQL_USERNAME:-blog} # blog 数据库用户名
readonly MYSQL_PASSWORD=${MYSQL_PASSWORD:-${PASSWORD}} # blog 数据库密码

# Redis 配置信息
readonly REDIS_HOST=${REDIS_HOST:-127.0.0.1} # Redis 主机地址
readonly REDIS_PORT=${REDIS_PORT:-6379} # Redis 监听端口
readonly REDIS_PASSWORD=${REDIS_PASSWORD:-${PASSWORD}} # Redis 密码

# 项目 配置
readonly DATA_DIR=${DATA_DIR:-/data/blog} # blog 各组件数据目录
readonly INSTALL_DIR=${INSTALL_DIR:-/opt/blog} # blog 安装文件存放目录
readonly CONFIG_DIR=${CONFIG_DIR:-/etc/blog} # blog 配置文件存放目录
readonly LOG_DIR=${LOG_DIR:-/var/log/blog} # blog 日志文件存放目录
readonly CA_FILE=${CA_FILE:-${CONFIG_DIR}/cert/ca.pem} # CA

# apiserver 配置
readonly APISERVER_HOST=${APISERVER_HOST:-127.0.0.1} # apiserver 部署机器 IP 地址
readonly APISERVER_INSECURE_HOST=${APISERVER_INSECURE_HOST:-127.0.0.1}
readonly APISERVER_INSECURE_PORT=${APISERVER_INSECURE_PORT:-8080}