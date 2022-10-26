package oms导入客户关系

/**
先用这个查询出org_id
="select id org_id, 0 is_default, 'admin' create_by , now() create_time,  '"&F2&"' customer_code from dc_oms.base_organization bo where nc_code ='"&M2&"';"


在拼接insert语句执行
INSERT INTO dc_oms.base_customer_org_relation (org_id, is_default, create_by, create_time, customer_code)
select id org_id, 0 is_default, 'admin' create_by , now() create_time,  'KH0004518' customer_code from dc_oms.base_organization bo where nc_code ='GYL024';




导出对应的大仓的关系
select bcor.customer_code "R1客户编码", bc.name "R1客户名称", bo.nc_code "nc编码", bo.name "对应大仓", bcor.org_id from dc_oms.base_customer_org_relation bcor
join dc_oms.base_organization bo on bo.id = bcor.org_id
join dc_oms.base_customer bc on bc.code = bcor.customer_code and bc.org_id = bcor.org_id
where length(bo.nc_code) >0


*/
