package main

import (
	"errors"
	"github.com/labstack/gommon/log"
)

type TenantService struct {
	err error
}

func GetTenantService() *TenantService {
	return new(TenantService)
}

func (t *TenantService) Err() error {
	return t.err
}

type Tenant struct {
	TenantId string
}

func (t *TenantService) ActiveTenant(userId int64) *Tenant {
	if t.err != nil {
		return nil
	}

	// 这里调用Tenant业务对象，给用户开启租户

	// TODO: 服务层错误不用处理，抛给外层，但是要记下关键信息
	//log.Error("这里记下关键信息，比如入参什么的，" +
	//	"不用记错误，错误可以抛给上层处理")

	return &Tenant{TenantId: "1"}
}

func (t *TenantService) InitBilling(tenant *Tenant) bool {
	if t.err != nil {
		return false
	}
	// 这里调用Billing业务对象给租户开启
	// 账单中心等等需要在开租户时就要有的功能模块

	// 注意 Service里记日志只
	t.err = errors.New("错误")

	return true
}

// 假设这是一个路由绑定的控制器方法
func ActivateTenantForUser() {
	tenantService := GetTenantService()
	tenant := tenantService.ActiveTenant(1)
	tenantService.InitBilling(tenant)

	if err := tenantService.Err(); err != nil {
		log.Error()
		// 返回错误响应给客户端
	}

	// 返回正常响应给客户端
}

func main() {
	ActivateTenantForUser()
}
