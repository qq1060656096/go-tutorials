local keysOkCount = 0;
local operationOk = nil;
local clientId = ARGV[1];
local expired = ARGV[2];
local tmpClientId = nil;
local lockNamesLen = table.getn(KEYS);
local setOkKeys = {}
-- 参数小于2参数错误
if (table.getn(ARGV) < 2) then
    return -1
end

for i, v in ipairs(KEYS) do
    operationOk  = redis.call('set', KEYS[i], clientId, 'PX', expired, 'NX');
    if (operationOk) then
        keysOkCount = keysOkCount + 1;
        setOkKeys[keysOkCount] = KEYS[i];
    end;
end;

-- 批量锁成功
if (keysOkCount == lockNamesLen) then
    return keysOkCount;
end;

-- 批量锁失败, 删除当前客户端已经设置的锁
for i, v in ipairs(setOkKeys) do
    tmpClientId  = redis.call('get', setOkKeys[i]);
    if (tmpClientId == clientId) then
        redis.call('del', setOkKeys[i]);
        keysOkCount = keysOkCount - 1;
    end;
end
return keysOkCount;
