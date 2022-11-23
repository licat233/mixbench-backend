#!/bin/bash

goctl api plugin -plugin goctl-swagger="swagger -filename ../static/swagger.json" -api ./admin.api -dir .
