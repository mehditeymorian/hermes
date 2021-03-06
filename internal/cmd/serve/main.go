package serve

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mehditeymorian/hermes/internal/config"
	"github.com/mehditeymorian/hermes/internal/db/mongo"
	"github.com/mehditeymorian/hermes/internal/db/store"
	"github.com/mehditeymorian/hermes/internal/emq"
	"github.com/mehditeymorian/hermes/internal/http/handler"
	"github.com/mehditeymorian/hermes/internal/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func Command(cfgFile string) *cobra.Command {
	serveCommand := &cobra.Command{ //nolint:exhaustivestruct
		Use:   "serve",
		Short: "signaling server",
		Run:   run,
	}

	return serveCommand
}

func run(cmd *cobra.Command, _ []string) {
	cfgFile := cmd.Flag("config").Value.String()

	cfg := config.Load(cfgFile)

	logger := log.New(cfg.Logger)

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	emqClient := emq.Connect(cfg.Emq)

	emqx := emq.Emq{Client: emqClient, Logger: logger}

	dbClient, err := mongo.Connect(cfg.DB)
	if err != nil {
		zap.L().Fatal("failed to connect to db", zap.Error(err))
	}

	dbStore := store.New(dbClient)

	app := fiber.New()

	app.Use(fiberLogger.New())

	handler.Room{
		Logger: logger,
		Store:  dbStore,
		Emq:    emqx,
	}.Register(app)

	zap.L().Fatal("failed to run app", zap.Error(app.Listen(":3000")))
}
