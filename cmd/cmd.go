package cmd

import (
	"context"
	"simplicity/env"
	"simplicity/llmbridge"
	"simplicity/server"
)

func Execute() {

	ctx := context.Background()
	llmclient := llmbridge.NewDefaultLLMClient(ctx, env.GetModel())

	srv := server.NewServer(ctx, llmclient)
	server.ManageRoutes(ctx, srv)

	srv.Start(env.GetServerPort())

}
