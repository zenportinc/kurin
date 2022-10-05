package main

import (
	"github.com/zenportinc/kurin"
	"github.com/zenportinc/kurin/example/adapters/http"
	"github.com/zenportinc/kurin/example/engine"
	"github.com/zenportinc/kurin/example/providers/example"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()

	// Providers
	exampleProviderFactory := example.NewFactory()

	// Engine
	e := engine.NewFactory(exampleProviderFactory).NewEngine()

	// App
	app := kurin.NewApp("Example", http.NewAdapter(e, 7272, logger))
	app.RegisterSystems(exampleProviderFactory)
	app.SetLogger(logger.Sugar())
	app.Run()
}
