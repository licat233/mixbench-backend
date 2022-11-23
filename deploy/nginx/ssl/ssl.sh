#!/bin/bash

#进入证书存放目录
cd ~/nginx/ssl/mixbench.net/

#------------------ 方法一 -----------------
#1.生成私钥(输入一个密码:123456)
openssl genrsa -des3 -out private.key 2048

#2.生成证书
#CN
#GuangDong
#GuangZhou
#IBM
#IBM
#mixbench.net
#licat233@gmail.com
openssl req -new -key private.key -out server.csr

#3.生成服务器的私钥，去除密钥口令(123456)
openssl rsa -in private.key -out server.key

#4.使用私钥为证书请求签名，生成给服务器签署的证书，格式是x509的PEM格式(mac密码)
sudo openssl x509 -req -in server.csr -out server.crt -signkey server.key -days 3650

#-------------------- 方法二 ------------------
# 生成PKCS12格式的密钥文件
keytool -genkey -alias panel.mixbench.net -keyalg RSA -keysize 2048 -sigalg SHA256withRSA -storetype PKCS12 -keystore server.p12 -dname CN=panel.mixbench.net,OU=Test,O=pkslow,L=Guangzhou,C=CN -validity 731 -storepass 123456 -keypass 123456

# 导出pem(certificate)
openssl pkcs12 -nokeys -in ./server.p12 -out server.pem

# 导出key
openssl pkcs12 -nocerts -nodes -in ./server.p12 -out server.key
