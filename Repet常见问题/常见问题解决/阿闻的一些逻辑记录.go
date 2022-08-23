package main

/**
// todo  上架会拉取库存，前提是redis不能有那个skuid的值，需要删除redis的skuId才回去拉取子龙库存








todo  mt自动上架逻辑
1： 查询渠道的仓库绑定
2： 查询渠道的第三方门店id
3： 循环查询到的仓库绑定关系，按照仓库处理
4： 查询价格表有价格的，没有上过架的
5： 判断库存(医疗互联网渠道不需要判断库存)
6： 上架表不存在的话，会新增快照，会修改快照
	上架表存在的话，查询快照信息，如果不存在新增快照信息，存在快照就修改
	 这里对于阿闻渠道同意都会插入hasStock表，显示有无库存
7： 格式化第三方的数据，调用第三方的同步商品的接口，调整上架状态，进行上架，也会把库存同步上去



todo 同步门店仓价格的
1：先是在渠道设置里面点击同步仓库价格，调用子龙的查询门店仓库的价格保存到 price_sync_response
2：每次取出100条数据进行处理 这一步是在： c.AddFunc("@every 60s", SyncProductPriceSync)
3：将处理后的数据插入到价格表中 price_sync 此时表中的数据 enable=1字段
4：查询门店仓价格表 enable = 1的，每次处理4000条数据，放到表PriceSyncRecord中, 四个渠道的每个渠道都保存一份，然后修改price_sync的 enable=0， 这一步是在 c.AddFunc("@every 60s", Pushmqgo)任务中执行
5：推送到mq，这里查询表PriceSyncRecord中is_push= 0微推送的。每次处理4000条，推送到四个渠道的mq(c.AddFunc("@every 60s", PushMqData))，消费是在product-center的四个任务中

6：消费门店仓价格：
//价格同步消费处理--阿闻
go SyncProductPriceAW()
//价格同步消费处理--美团
go SyncProductPriceMT()
//价格同步消费处理--饿了么
go SyncProductPriceElm()
//价格同步消费处理--京东
go SyncProductPriceJD()


todo 前置仓价格同步(前置仓价格是导入进去的)
1： 任务 处理批量导入前置仓价格信息 插入到定时任务中，task-content = 19
2： 处理任务 c.AddFunc("@every 5s", pubilchProductTask(batchHandleQzcPrice, services.DealAsyncTaskListByQzcPrice(19), 30))
3： 插入或者更新到表qzc_price_sync中
4： 更新价格的时候会保存到 qzc_price_sync_record中，里面记录了用户信息，插入的时候不会保存，只会向qzc_price_sync中插入数据
5： 这里直接是四个渠道的定时任务更新渠道价格（ todo 前置仓价格是没有走mq的注意点）
// 每晚0点更新前一天有更新过的前置仓价格
c.AddFunc("0 0 * * *", SyncQzcPriceAwen) // 阿闻
c.AddFunc("0 0 * * *", SyncDsPriceAwen)  // 阿闻电商仓价格同步
c.AddFunc("0 0 * * *", SyncQzcPriceMt)   // 美团
c.AddFunc("0 0 * * *", SyncQzcPriceElm)  // 饿了么
c.AddFunc("0 0 * * *", SyncQzcPriceJddj) // 京东



//
*/
