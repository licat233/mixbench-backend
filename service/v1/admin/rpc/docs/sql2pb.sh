#!/bin/bash
###
 # @Author: licat
 # @Date: 2023-01-11 15:40:07
 # @LastEditors: licat
 # @LastEditTime: 2023-01-15 14:30:06
 # @Description: licat233@gmail.com
### 

# sql2pb -go_package ./pb -host localhost -package pb -password root -port 3306 -schema mixbench -service_name admin -user root > pb.proto
sql2pb -field_style sql_pb -ignore_columns create_at,update_at,is_deleted,delete_at -go_package ./pb -host localhost -package pb -password "" -port 3306 -schema mixbench -service_name admin -user root > mixbench.proto



