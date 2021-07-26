package main

//import (
//	"github.com/golang/glog"
//	"github.com/spf13/cast"
//	"strings"
//)

// 查询效期列表数据
//func (this Visual) SearchEffectiveList(ctx context.Context, in *sv.EffectiveManagementResVo) (*sv.EffectiveManagementResponseData, error) {
//	glog.Info("SearchEffectiveList: ", rpkit.JsonEncode(in))
//
//	response := sv.EffectiveManagementResponseData{}
//
//	conn := NewDbConn()
//
//	sqlCount := `select count(*) totalcount from sc_stock.effective_management a join  dc_dispatch.warehouse b on a.warehouse_id=b.id where 1= 1`
//	if len(in.WarehouseId) > 0 {
//		sqlCount += " and warehouse_id in (" + in.WarehouseId + ")"
//	}
//	if len(in.SkuId) > 0 {
//		sqlCount += " and a.sku_id = '" + in.SkuId + "'"
//	}
//	if len(in.ThirdSkuId) > 0 {
//		sqlCount += " and a.third_sku_id = '" + in.ThirdSkuId + "'"
//	}
//	if in.Status > 0 {
//		sqlCount += " and a.effective_state = " + cast.ToString(in.Status)
//	}
//	if in.EndDay > 0 || in.StartDay > 0 {
//		sqlCount += " and a.effective_days > " + cast.ToString(in.StartDay) + " and a.effective_days < " + cast.ToString(in.EndDay)
//	}
//
//	conn.ShowSQL(true)
//	managements := make([]*sv.EffectiveManagementResponse, 0)
//	var totalcount int32
//	conn.SQL(sqlCount).Get(&totalcount)
//
//
//	sqlString := strings.Replace(sqlCount, "count(*) totalcount", "a.*, b.name warehouse_name", 1)
//	sqlString += " limit ?,?"
//	err := conn.SQL(sqlString, (in.PageIndex-1)*in.PageSize, in.PageSize).Find(&managements)
//
//	if err != nil {
//		glog.Error(" table.FindAndCount(&managements)", err.Error())
//		return &response, err
//	}
//
//	response.Data = managements
//	response.Total = totalcount
//	return &response, nil
//
//}
//
