//go:build wireinject
// +build wireinject

package Route

import (
	"ct-backend/Controller"
	"ct-backend/Middleware"
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
		Services.JwtServiceProvider,

		wire.Bind(new(Controller.IAuthController), new(*Controller.AuthController)),
		wire.Bind(new(Services.IAuthService), new(*Services.AuthService)),
		wire.Bind(new(Repository.IAuthRepository), new(*Repository.AuthRepository)),
		wire.Bind(new(Services.IJwtService), new(*Services.JwtService)),
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

func CommonMiddlewareDI() *Middleware.CommonMiddleware {
	panic(wire.Build(wire.NewSet(
		Middleware.CommonMiddlewareProvider,
		Services.JwtServiceProvider,

		wire.Bind(new(Services.IJwtService), new(*Services.JwtService)),
	),
	))
	return &Middleware.CommonMiddleware{}
}

func PurchaseDI(db *gorm.DB) *Controller.PurchaseController {
	panic(wire.Build(wire.NewSet(
		Repository.ProductRepositoryProvider,
		Repository.PurchaseRepositoryProvider,
		Services.PurchaseServiceProvider,
		Controller.PurchaseControllerProvider,

		wire.Bind(new(Controller.IPurchaseController), new(*Controller.PurchaseController)),
		wire.Bind(new(Services.IPurchaseService), new(*Services.PurchaseService)),
		wire.Bind(new(Repository.IProductRepository), new(*Repository.ProductRepository)),
		wire.Bind(new(Repository.IPurchaseRepository), new(*Repository.PurchaseRepository)),
	),
	))
	return &Controller.PurchaseController{}
}

func InvoiceDI(db *gorm.DB) *Controller.InvoiceController {
	panic(wire.Build(wire.NewSet(
		Repository.InvoiceRepositoryProvider,
		Services.InvoiceServiceProvider,
		Controller.InvoiceControllerProvider,

		wire.Bind(new(Controller.IInvoiceController), new(*Controller.InvoiceController)),
		wire.Bind(new(Services.IInvoiceService), new(*Services.InvoiceService)),
		wire.Bind(new(Repository.IInvoiceRepository), new(*Repository.InvoiceRepository)),
	),
	))
	return &Controller.InvoiceController{}
}
