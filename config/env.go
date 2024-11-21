package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	// ==================================================== ROUTING ==================================================== //

	WebPublicPrefix = GetEnv("WEB_PUBLIC_PREFIX", "/public")
	SERVERPort      = GetEnv("BE_SERVER_PORT", ":9070")
	ApiPrefix       = GetEnv("BE_API_PREFIX", "/api")
	WebHost         = GetEnv("BE_WEB_HOST", "http://localhost:9071")
	APIHost         = "http://localhost" + SERVERPort
	AppPrefix       = GetEnv("BE_BASE_URL", "/tkbai")
	JwtKey          = GetEnv("BE_SV_JWT_KEY", "LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U=")
	DbUrl           = GetEnv("DB_URL", "tkbai:rytdin-ryqriN-kohqi6@tcp(127.0.0.1:3306)/tkbai?parseTime=true")

	// ==================================================== SESSION ==================================================== //

	SessionCookieName = GetEnv("SESSION_NAME", "TKBAI-SESSION")
	AppSessionSecret  = GetEnv("SESSION_SECRET", "LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U=")

	// ==================================================== NONCE ==================================================== //

	StyleSrcNonce  = GetEnv("STYLE_NONCE", "K4qfk2XrYB6uE81e")
	ScriptSrcNonce = GetEnv("SCRIPT_NONCE", "L4Mcme5VgMor9KF0")

	// ==================================================== CSRF ==================================================== //

	CsrfToken = GetEnv("CSRF_TOKEN", "1nLHl1Msf9cn0k+SPE1c0UJp1IxA6uH1jsxS2vnK")
)

func GetEnv(key, fallback string) (value string) {
	logger := Log
	if value, ok := os.LookupEnv(key); ok {
		logger.Debug().Str(key, value).Msg("Env")
		return value
	}
	logger.Error().Str(key, fallback).Msg("Fallback")
	return fallback
}
