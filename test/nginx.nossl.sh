#!/bin/bash
# 请先进入 nginx 目录 启动服务 docker-compose up
curl -v 'http://127.0.0.1:7903/lnks_homecloud.clh021.webshell.ShellService/Run2' \ 
  -H 'x-grpc-web: 1' \
  -H 'content-type: application/grpc-web+proto' \
  --data-binary @data.bin \
  --compressed \
  --output output2.log
# data.bin,output 输出请使用 cat 查看内容