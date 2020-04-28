local clientId = ARGV[1];
local tmpClientId = nil;
local keysDelOkCount = 0;

for i, v in ipairs(KEYS) do
    tmpClientId  = redis.call('get', KEYS[i]);
    if (tmpClientId == clientId) then
        keysDelOkCount = keysDelOkCount + redis.call('del', KEYS[i]);
    end;
end

return keysDelOkCount;
