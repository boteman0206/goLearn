package redis相关

/**

scriptStr := `
	local delimiter = "-"
 	local skuArray = {}
 	local ckArray = {}
 	local stockArray = {}
 	local redisStockArray = {}
 	local warehouseId = ARGV[2]
	local noStockStr = ""
	local skuStr = {}
	for i= 1, ARGV[1] do
		local sourceStr = KEYS[i]
		local findStart,findEnd = string.find(KEYS[i],delimiter)
		local stock_key = string.sub(sourceStr, 1, findStart-1)
		skuArray[i]=stock_key
		local stock_value = string.sub(sourceStr, findStart+1, string.len(KEYS[i]))
		stockArray[i]=stock_value
		-- 查询redis库存信息
		local stock_redis_value = tonumber(redis.call('HGET', stock_key, warehouseId))
		if stock_redis_value == nil then
			skuStr[1] = stock_key
			return skuStr
		end
		redisStockArray[i]=stock_redis_value
		-- 判断库存是否满足
		if(stock_redis_value-stock_value >= 0) then
			if (stock_redis_value-stock_value==0) then
				noStockStr = noStockStr .. "," .. stock_key
			end
			redisStockArray[i]=stock_redis_value
		else
			-- 如果不满足，则直接return出去
			skuStr[1] = stock_key
			return skuStr
		end
	 end
	 -- 验证通过后去redis操作数据
	 for i= 1, ARGV[1] do
		-- 查询redis库存信息
		local stock_redis_value = tonumber(redis.call('HGET', skuArray[i], warehouseId))
		local stock_value=tonumber(stockArray[i])
		-- 扣减库存
		local stock_num=stock_redis_value-stock_value
		-- 扣减完成后重新设置redis
		redis.call('HSET',skuArray[i],warehouseId,stock_num)
	 end
	 return "1|" .. noStockStr
	`
//处理lua参数格式
result := redisConn.Eval(scriptStr, paramsStrArr, len(paramsStrArr), WarehouseId)
resultType := fmt.Sprintf("%T", result.Val())



todo :  paramsStrArr  这里的KEYS取出的就是就是第一个参数的值, 看源码Eval的KEYS的长度就是取得这个参数的在里面执行了len()方法
	   len(paramsStrArr), WarehouseId 这两个代表的是ARG[1],ARG[2]
*/
