package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var (
	// ==================================================== ROUTING ==================================================== //

	WebPublicPrefix = GetEnv("WEB_PUBLIC_PREFIX", "/public")
	SERVERPort      = GetEnv("BE_SERVER_PORT", ":9070")
	ApiPrefix       = GetEnv("BE_API_PREFIX", "/api")
	WebHost         = GetEnv("BE_WEB_HOST", "http://localhost:9071")
	APIHost         = "http://localhost" + SERVERPort
	AppPrefix       = GetEnv("BE_BASE_URL", "/tkbai")
	AdminLoginURL   = WebHost + AppPrefix + "/login/admin"
	JwtKey          = GetEnv("BE_SV_JWT_KEY", "LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U=")

	// ==================================================== SESSION ==================================================== //

	SessionCookieName = GetEnv("SESSION_NAME", "TKBAI-SESSION")
	AppSessionSecret  = GetEnv("SESSION_SECRET", "LmPZJbddZ9uXW4JE7g6N9R8ZdmDRv5vYihZJRBcOz7U=")

	// ==================================================== NONCE ==================================================== //

	StyleSrcNonce  = GetEnv("STYLE_NONCE", "K4qfk2XrYB6uE81e")
	ScriptSrcNonce = GetEnv("SCRIPT_NONCE", "L4Mcme5VgMor9KF0")

	// ==================================================== IAM ==================================================== //

	IAMURL                  = GetEnv("BE_KC_URL", "http://localhost:8080")
	IAMDockerURL            = GetEnv("BE_KC_DOCKER_URL", "http://keycloak:8080")
	IAMClientSecret         = GetEnv("BE_KC_SECRET", "rK4sSxVGdIjnEKnyBFzdymlN62stQ72m")
	IAMClientID             = GetEnv("BE_KC_ID", "tkbai")
	IAMRealm                = GetEnv("BE_KC_REALM", "tkbai_dev")
	IAMState                = GetEnv("BE_KC_STATE", "authExt")
	IAMConfigURL            = IAMURL + "/realms/" + IAMRealm
	IAMDockerConfigURL      = IAMDockerURL + "/realms/" + IAMRealm
	IAMLoginRedirectPath    = GetEnv("BE_KC_LOGIN_REDIRECT_PATH", "")
	IAMLoginRedirect302Path = GetEnv("BE_KC_LOGIN_302REDIRECT_PATH", "")
	IAMLogoutRedirectPath   = GetEnv("BE_KC_LOGOUT_REDIRECT_PATH", "")
	IAMLoginRedirectURL     = APIHost + AppPrefix + IAMLoginRedirectPath
	IAMLoginRedirect302URL  = WebHost + AppPrefix + IAMLoginRedirect302Path
	IAMLogoutRedirectURL    = APIHost + AppPrefix + IAMLogoutRedirectPath
	TkbaiDB                 = GetEnv("BE_TKBAI_DB_URL", "root:03IZmt7eRMukIHdoZahl@tcp(mysql:3306)/tkbai")
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
