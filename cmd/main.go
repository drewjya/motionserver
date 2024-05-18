package main

import (
	"motionserver/app/middleware"
	"motionserver/app/module/auth"
	"motionserver/app/module/cart"
	"motionserver/app/module/category"
	"motionserver/app/module/compro"
	"motionserver/app/module/gallery"
	"motionserver/app/module/news"
	"motionserver/app/module/product"
	"motionserver/app/router"
	"motionserver/internal/bootstrap"
	"motionserver/internal/bootstrap/database"
	"motionserver/internal/bootstrap/minio"
	"motionserver/utils/config"

	fxzerolog "github.com/efectn/fx-zerolog"
	"go.uber.org/fx"
)

// @title                       Go Fiber Starter API Documentation
// @version                     1.0
// @description                 This is a sample API documentation.
// @termsOfService              http://swagger.io/terms/
// @contact.name                Developer
// @contact.email               bangadam.dev@gmail.com
// @license.name                Apache 2.0
// @license.url                 http://www.apache.org/licenses/LICENSE-2.0.html
// @host                        localhost:8080
// @schemes                     http https
// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description                 "Type 'Bearer {TOKEN}' to correctly set the API Key"
// @BasePath                    /
func main() {
	fx.New(
		/* provide patterns */
		// config
		fx.Provide(config.NewConfig),
		// logging
		fx.Provide(bootstrap.NewLogger),
		// fiber
		fx.Provide(bootstrap.NewFiber),
		// database
		fx.Provide(database.NewDatabase),
		// minio
		fx.Provide(minio.NewMinio),

		// middleware
		fx.Provide(middleware.NewMiddleware),
		// router
		fx.Provide(router.NewRouter),

		// provide modules

		auth.NewAuthModule,
		category.NewCategoryModule,
		product.NewProductModule,
		gallery.NewGalleryModule,
		cart.NewCartModule,
		news.NewNewsModule,
		compro.NewComproModule,

		// start aplication
		fx.Invoke(bootstrap.Start),

		// define logger
		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
