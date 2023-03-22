package netw

import "github.com/xpp777/game/iface"

// BaseRouter 实现router时，先嵌入这个基类，然后根据需要对这个基类的方法进行重写
type BaseRouter struct{}

// 这里之所以BaseRouter的方法都为空，
// 是因为有的Router不希望有PreHandle或PostHandle
// 所以Router全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化

// PreHandle -
func (br *BaseRouter) PreHandle(req iface.Request) {}

// Handle -
func (br *BaseRouter) Handle(req iface.Request) {}

// PostHandle -
func (br *BaseRouter) PostHandle(req iface.Request) {}
