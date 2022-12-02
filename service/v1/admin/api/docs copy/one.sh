#!/bin/bash

goctl api go -api admin.api -dir=../ -style goZero
goctl api plugin -plugin goctl-swagger="swagger -filename ../static/swagger.json" -api ./admin.api -dir .
goctl api ts -api admin.api -dir ../static
