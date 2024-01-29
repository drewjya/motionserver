package bootstrap

import (
	"context"
	"flag"
	"template_clean_architecture/app/middleware"
	"template_clean_architecture/app/router"
	"template_clean_architecture/internal/bootstrap/database"
	"template_clean_architecture/utils/config"
	"template_clean_architecture/utils/response"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func NewFiber(cfg *config.Config) *fiber.App {
	app := fiber.New(
		fiber.Config{
			AppName:           cfg.App.Name,
			EnablePrintRoutes: cfg.App.PrintRoutes,
			ErrorHandler:      response.ErrorHandler,
		},
	)
	response.IsProduction = cfg.App.Production
	return app
}

func Start(lifecycle fx.Lifecycle, cfg *config.Config, fiber *fiber.App, router *router.Router, middlewares *middleware.Middleware, database *database.Database, log zerolog.Logger) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				// Register middlewares & routes
				middlewares.Register()
				router.Register()

				// Custom Startup Messages
				host, port := config.ParseAddress(cfg.App.Port)
				if host == "" {
					if fiber.Config().Network == "tcp6" {
						host = "[::1]"
					} else {
						host = "0.0.0.0"
					}
				}

				// Information message
				log.Info().Msg(fiber.Config().AppName + " is running at the moment!")

				// Debug informations
				if !cfg.App.Production {
					prefork := "Enabled"
					log.Debug().Msgf("Version: %s", "-")
					log.Debug().Msgf("Host: %s", host)
					log.Debug().Msgf("Port: %s", port)
					log.Debug().Msgf("Prefork: %s", prefork)
					log.Debug().Msgf("Handlers: %d", fiber.HandlersCount())

				}

				// Listen the app (with TLS Support)

				go func() {
					if err := fiber.Listen(cfg.App.Port); err != nil {
						log.Error().Err(err).Msg("An unknown error occurred when to run server!")
					}
				}()

				database.ConnectDatabase()

				migrate := flag.Bool("migrate", false, "migrate the database")
				seeder := flag.Bool("seed", false, "seed the database")
				reset := flag.Bool("reset", false, "reset the database")
				flag.Parse()
				if *reset {
					database.ResetModels()
				} else if *migrate {
					database.MigrateModels()
				} else if *seeder {
					database.SeedModels()
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info().Msg("Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic().Err(err).Msg("")
				}

				log.Info().Msg("Running cleanup tasks...")
				log.Info().Msg("1- Shutdown the database")
				database.ShutdownDatabase()
				log.Info().Msgf("%s was successful shutdown.", cfg.App.Name)
				log.Info().Msg("\u001b[96msee you again👋\u001b[0m")

				return nil
			},
		},
	)
}
