#!/bin/bash

goctl api go -api captcha.api -dir=../ -style goZero
goctl api plugin -plugin goctl-swagger="swagger -filename ../static/swagger.json" -api captcha.api -dir .
goctl api ts -api captcha.api -dir ../static