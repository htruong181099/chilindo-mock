package routes

import "chilindo/src/product-service/controllers"

type IProductRoute interface {
	SetRouter()
}

type ProductRoute struct {
	ProductController controllers.IProductController
}

func (p ProductRoute) SetRouter() {
	//TODO implement me
	panic("implement me")
}
