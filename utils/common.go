package utils

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var EnvVars, _ = godotenv.Read(".env")

var Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
