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
	engineFactory := engine.NewFactory(exampleProviderFactory)
	e := engineFactory.NewEngine()

	// App
	a := kurin.NewApp("Example", http.NewHTTPAdapter(e, 7272, logger))
	a.RegisterSystems(exampleProviderFactory)
	a.Run()
}
