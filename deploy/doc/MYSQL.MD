# MYSQL笔记  

## 多表关联查询遇到重名字段时，写法如下  

```sql
SELECT
 l.uuid AS luckerUUID,
 pp.uuid AS prizespoolUUID,
 a.uuid AS activityUUID,
 p.uuid AS prizeUUID 
FROM
 `lucker` AS l
 LEFT JOIN `prizes_pool` AS pp ON l.prizes_pool_id = pp.id
 LEFT JOIN `activity` AS a ON pp.activity_id = a.id
 LEFT JOIN `prize` AS p ON pp.prize_id = p.id
WHERE
 l.is_deleted = 0 
 AND pp.is_deleted = 0 
 AND a.is_deleted = 0 
 AND p.is_deleted = 0 
 AND l.user_id = 1 
 AND pp.activity_id = 1;
```
