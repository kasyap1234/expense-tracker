package routes

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
	"log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/kasyap1234/expense-tracker/config"
	"github.com/kasyap1234/expense-tracker/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)
func init(){
	godotenv.Load()

}

func InitializeOAuth() *oauth2.Config {
	
	
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URI"),
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
    oauthConfig := InitializeOAuth()
    url := oauthConfig.AuthCodeURL("state")
    log.Printf("Redirecting to Google: %s", url)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}


func OauthCallback(w http.ResponseWriter, r *http.Request) {
	oauthConfig := InitializeOAuth()
	// code in the url from oauthcallback 

	code := r.URL.Query().Get("code")
	// exchange code for token ; 
	token, err := oauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}
	
	client := oauthConfig.Client(r.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	json.NewDecoder(resp.Body).Decode(&userInfo)

	var user models.User
	config.DB.FirstOrCreate(&user, models.User{
		Email:    userInfo.Email,
		Username: userInfo.Name,
	})

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
	frontendURL := "http://localhost:3000/dashboard?token=" + tokenString
    http.Redirect(w, r, frontendURL, http.StatusTemporaryRedirect)
	response :=map[string]string{
		"message": "Login Successful",
		"token": tokenString,
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response); 
	
}

func generateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
