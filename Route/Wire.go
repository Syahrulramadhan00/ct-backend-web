//go:build wireinject
// +build wireinject

package Route

import (
	"ct-backend/Controller"
	"ct-backend/Repository"
	"ct-backend/Services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func AuthDI(db *gorm.DB) *Controller.AuthController {
	panic(wire.Build(wire.NewSet(
		Repository.AuthRepositoryProvider,
		Services.AuthServiceProvider,
		Controller.AuthControllerProvider,

		wire.Bind(new(Controller.IAuthController), new(*Controller.AuthController)),
		wire.Bind(new(Services.IAuthService), new(*Services.AuthService)),
		wire.Bind(new(Repository.IAuthRepository), new(*Repository.AuthRepository)),
	),
	))
	return &Controller.AuthController{}
}

func ProductDI(db *gorm.DB) *Controller.ProductController {
	panic(wire.Build(wire.NewSet(
		Repository.ProductRepositoryProvider,
		Services.ProductServiceProvider,
		Controller.ProductControllerProvider,

		wire.Bind(new(Controller.IProductController), new(*Controller.ProductController)),
		wire.Bind(new(Services.IProductService), new(*Services.ProductService)),
		wire.Bind(new(Repository.IProductRepository), new(*Repository.ProductRepository)),
	),
	))
	return &Controller.ProductController{}
}
