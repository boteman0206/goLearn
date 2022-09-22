package oms导入客户关系

/**
先用这个查询出org_id
="select id org_id, 0 is_default, 'admin' create_by , now() create_time,  '"&F2&"' customer_code from dc_oms.base_organization bo where nc_code ='"&M2&"';"


在拼接insert语句执行
INSERT INTO dc_oms.base_customer_org_relation (org_id, is_default, create_by, create_time, customer_code)
select id org_id, 0 is_default, 'admin' create_by , now() create_time,  'KH0004518' customer_code from dc_oms.base_organization bo where nc_code ='GYL024';
*/
