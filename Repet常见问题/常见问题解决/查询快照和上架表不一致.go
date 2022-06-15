package main

/**

	select * from datacenter.store where name regexp '宠颐生海口远大分院';

	select * from dc_product.channel_store_product where finance_code='CX0373' and product_id=1017168 and channel_id=1;

	select * from dc_product.channel_store_product where id=24518716;
	-- 25852628
	select * from dc_product.channel_product_snapshot where id=25818116;


	update dc_product.channel_store_product set snapshot_id=25818116, update_date= now() where id=24518716;



	-- todo 修复快照数据和上架表的id不一致的数据

 	update dc_product.channel_store_product a
     set a.snapshot_id = (select b.id from dc_product.channel_product_snapshot b
                        where  a.finance_code=b.finance_code
                          and a.product_id=b.product_id and a.channel_id=b.channel_id),
                 update_date=now() where id in(
           select * from (select id from dc_product.channel_store_product
           where finance_code='CX0011' and product_id in(1023611) and channel_id=1) c1  -- 6927602
         )


*/
