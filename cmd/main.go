package main

import (
	"template_clean_architecture/app/middleware"
	"template_clean_architecture/app/module/auth"
	"template_clean_architecture/app/router"
	"template_clean_architecture/internal/bootstrap"
	"template_clean_architecture/internal/bootstrap/database"
	"template_clean_architecture/utils/config"

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
		// middleware
		fx.Provide(middleware.NewMiddleware),
		// router
		fx.Provide(router.NewRouter),

		// provide modules

		auth.NewAuthModule,

		// start aplication
		fx.Invoke(bootstrap.Start),

		// define logger
		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
