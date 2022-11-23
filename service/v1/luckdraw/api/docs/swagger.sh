#!/bin/bash

# 生成# swagger文档
goctl api plugin -plugin goctl-swagger="swagger -filename ../static/swagger.json" -api luckdraw.api -dir .
