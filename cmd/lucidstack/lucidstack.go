package main

import (
	"github.com/lucidstackhq/lucidstack/internal/app/lucidstack"
	"github.com/lucidstackhq/lucidstack/internal/pkg/env"
)

func main() {
	lucidstack.NewServer(&lucidstack.ServerConfig{
		Host:          env.GetOrDefault("HOST", "0.0.0.0"),
		Port:          env.GetOrDefault("PORT", "8000"),
		MongoEndpoint: env.GetOrDefault("MONGO_ENDPOINT", "mongodb://localhost:27017"),
		MongoDatabase: env.GetOrDefault("MONGO_DB", "lucidstack"),
		JwtSigningKey: env.GetOrDefault("JWT_KEY", "secret"),
	}).Start()
}
