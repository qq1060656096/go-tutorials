## 1. redis lua脚本

```
# KEYS 长度
eval "return table.getn(KEYS)" 2 key1 key2
# (integer) 2

# ARGV 长度
eval "return table.getn(ARGV)" 2 key1 key2 arg1 arg2 arg3
(integer) 3
```


```
# 分布式锁: 加锁
redis-cli -h 199.199.199.199 -p 16379 -a 000000 --eval lua/redis.lock.lua key1 key2  keyn, clientId expired
redis-cli -h 199.199.199.199 -p 16379 -a 000000 --eval lua/redis.lock.lua key1 key2  keyn, 客户端id 过期时间(毫秒级)
redis-cli -h 199.199.199.199 -p 16379 -a 000000 --eval lua/redis.lock.lua key1 key2 , clientId1 200000

# 分布式锁: 解锁
redis-cli -h 199.199.199.199 -p 16379 -a 000000 --eval lua/redis.unlock.lua key1 key2  keyn, 客户端id 过期时间(毫秒级)
redis-cli -h 199.199.199.199 -p 16379 -a 000000 --eval lua/redis.unlock.lua key1 key2 , clientId1
```