package 常见问题解决




// 重新推送mq查询订单数据
SELECT trim(both '"' from json_extract(content,'$.order_id')) as order_id
from dc_order.Mq_info where quene='dc-sz-qqd-dispatch' and lastdate>='2021-11-12 00:00:00' ;



SELECT trim(both '"' from json_extract(content,'$.order_id')) as order_id
from dc_order.Mq_info where quene='dc-sz-qqd-dispatch' and lastdate>='2021-11-12 00:00:00' ;



