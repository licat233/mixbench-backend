-- --------------
-- 创建 admin 数据库
-- --------------
CREATE DATABASE IF NOT EXISTS mixbench_admin;

-- --------------
-- 创建 luckydraw 数据库
-- --------------
CREATE DATABASE IF NOT EXISTS mixbench_luckydraw;

-- --------------
-- 创建 mall 数据库
-- --------------
CREATE DATABASE IF NOT EXISTS mixbench_mall;

----------------
-- Mysql创建mixbench用户
----------------
-- create user 'mixbench'@'%' identified by 'yI2~qF3#wY3$';

---------------------
-- 将mixbench用户绑定到mixbench数据库上
---------------------
-- grant all privileges on mixbench.* to 'mixbench'@'%';

------------------
-- 赋予增删改查的权限
------------------
-- grant select,delete,update,create,drop on mixbench.* to 'mixbench'@'%';

----------
-- 刷新权限
----------
-- flush privileges;
