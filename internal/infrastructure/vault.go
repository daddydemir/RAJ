package infrastructure

import (
	"context"
	"log/slog"
	"os"

	"github.com/hashicorp/vault/api"
)

func NewVaultClient() *api.Client {
	vaultAddr := os.Getenv("VAULT_ADDR")
	if vaultAddr == "" {
		vaultAddr = "http://localhost:8200"
	}

	config := &api.Config{
		Address: vaultAddr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		slog.Error("api.NewClient(config)", "message", err.Error())
		panic(err)
	}

	vaultToken := os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		slog.Error("os.Getenv(VAULT_TOKEN)", "message", "VAULT_TOKEN is not set")
		panic("VAULT_TOKEN is not set")
	}

	client.SetToken(vaultToken)

	return client
}

func GetSecrets(client *api.Client) map[string]interface{} {

	secret, err := client.KVv2("daddydemir").Get(context.Background(), "raj")
	if err != nil {
		slog.Error("client.KVv2(daddydemir).Get(ctx, raj)", "message", err.Error())
		panic(err)
	}

	return secret.Data
}

func WriteSecrets(client *api.Client, path string, data map[string]interface{}) {
	_, err := client.KVv2("daddydemir").Put(context.Background(), "raj", data)
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
}
