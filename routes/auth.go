package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"github.com/kasyap1234/expense-tracker/config"
	"github.com/kasyap1234/expense-tracker/models"
	"golang.org/x/oauth2/google"
)

// google ouath2 configuration
var oauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	RedirectURL:  "https://localhost:8000/auth/google/callback",
	Scopes:       []string{"openid", "email", "profile"},
	Endpoint:     google.Endpoint,
}

// gogole login handler : redirection
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

// google oauth2 callback handler : gets code and exchanges for token ;
func OauthCallback(w http.ResponseWriter, r *http.Request) {
	// get code from url query parameters
	code := r.URL.Query().Get("code")
	// exchange code for token ;
	token, err := oauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}
	// get user info from google api using token
	client := oauthConfig.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// decode user info from response body
	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	json.NewDecoder(resp.Body).Decode(&userInfo)
	var user models.User

	// check if user already exists in database if not create the new user or update it
	config.DB.FirstOrCreate(&user, models.User{
		Email:    userInfo.Email,
		Username: userInfo.Name,
	})
	// generate jwt token and set it as cookie for persistence ;
	tokenString, err := generateToken(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
	w.Write([]byte("Login successful"))

}

func generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
