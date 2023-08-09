package auth

import (
	"account-app/internal/initializers"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
)

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

func Init(config initializers.EnvironmentConfig) *Auth {
	return &Auth{
		Issuer:        config.JWT.JWTIssuer,
		Audience:      config.JWT.JWTAudience,
		Secret:        config.JWT.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour - 24,
		CookieDomain:  config.JWT.CookieDomain,
		CookiePath:    "/",
		CookieName:    "_HOST-refresh-token",
	}
}

type JwtUser struct {
	ID          string `json:"id"`
	UserName    string `json:"user_name"`
	UserEmail   string `json:"user_email"`
	AccountType string `json:"account_type"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokenPair(user *JwtUser) (TokenPairs, error) {
	//Create token
	token := jwt.New(jwt.SigningMethodHS256)
	//Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.UserName
	claims["email"] = user.UserEmail
	claims["acc_type"] = user.AccountType
	claims["sub"] = fmt.Sprint(user.ID)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	// Set the expiry for JWT
	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	//Create a refresh token and set claims
	signedAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()

	//set the expiry for the refresh token
	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	//create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}
	//create token pairs and populate with signed tokens
	var tokenPairs = TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	//return token pairs
	return tokenPairs, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    refreshToken,
		Expires:  time.Now().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   j.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func (j *Auth) GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   j.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}
