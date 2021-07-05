package main

import "fmt"

// 订单记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// 订单的详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

//抽象模式的工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// 为关系型数据库的orderMainDAO实现
type RDBMainDAO struct{}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Println("rbd main save")
}

//RDBDetailDAO 为关系型数据库的OrderDetailDAO实现
type RDBDetailDAO struct{}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

//RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

//XMLMainDAO XML存储
type XMLMainDAO struct{}

//SaveOrderMain ...
func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

//XMLDetailDAO XML存储
type XMLDetailDAO struct{}

// SaveOrderDetail ...
func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

//XMLDAOFactory 是RDB 抽象工厂实现
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}

// 抽象工厂方法模式

func main() {

	factory := XMLDAOFactory{}
	dao := factory.CreateOrderMainDAO()
	dao.SaveOrderMain()

	rdbdaoFactory := RDBDAOFactory{}
	mainDAO := rdbdaoFactory.CreateOrderMainDAO()
	mainDAO.SaveOrderMain()

}

/**


首先来看看两者的定义区别：
工厂模式 定义一个用于创建对象的接口，让子类决定实例化哪一个类
抽象工厂模式 为创建一组相关或相互依赖的对象提供一个接口，而且无需指定他们的具体类


抽象工厂模式是所有形式的工厂模式中最为抽象和最具一般性的一种形态。抽象工厂模式与工厂方法模式最大的区别在于，
工厂方法模式针对的是一个产品等级结构，
抽象工厂模式则需要面对多个产品等级结构。


todo https://design-patterns.readthedocs.io/zh_CN/latest/creational_patterns/builder.html
*/
