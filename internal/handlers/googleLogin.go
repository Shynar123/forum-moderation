package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"forum/internal/cookies"
	"forum/internal/types"
)

const (
	googleAuthURL     = "https://accounts.google.com/o/oauth2/auth"
	googleTokenURL    = "https://accounts.google.com/o/oauth2/token"
	googleUserInfoURL = "https://www.googleapis.com/oauth2/v3/userinfo"
)

var config = &types.Config{
	ClientID:     "318024681907-ld37vf67i4gt7desbhjh4niqai5adtcl.apps.googleusercontent.com", // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
	ClientSecret: "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",                                      // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
	Endpoint:     "http://localhost:8000/google/callback",
	Scopes:       "email profile",
}

type googleInfo struct {
	Sub      string `json:"sub"`
	Username string `json:"name"`
	Email    string `json:"email"`
}

func (h *Handler) handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}
	url := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s", googleAuthURL, config.ClientID, config.Endpoint, config.Scopes)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	data := strings.NewReader(fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code", code, config.ClientID, config.ClientSecret, config.Endpoint))
	resp, err := http.Post(googleTokenURL, "application/x-www-form-urlencoded", data)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
	
		ErrorPage(w, r, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)

		ErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	token := h.GetToken(body)


	req, err := http.NewRequest("GET", googleUserInfoURL, nil)
	if err != nil {
		fmt.Println("Google get:", err)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Client err:", err)
		return
	}
	defer response.Body.Close()

	info, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Reading response ", err)
		return
	}
	var googleUser googleInfo
	err = json.Unmarshal(info, &googleUser)
	if err != nil {
		fmt.Println("Unmarshal error ", err)
		return
	}
	user := &types.CreateUserData{
		Username: strings.ReplaceAll(googleUser.Username, " ", ""),
		Email:    googleUser.Email,
		Password: googleUser.Sub,
	}
	existBool, _ := h.service.UserService.CheckUserExists(user)

	if !existBool {
		err := h.service.UserService.CreateUser(user)
		if err != nil {
			fmt.Println("Create user error:", err)
		}
	}

	userid := h.service.UserService.GetLoginId(user)
	cookieToken := cookies.SetCookie(w)
	h.service.UserService.AddToken(userid, cookieToken)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handler) GetToken(body []byte) string {
	var temp map[string]interface{}
	err := json.Unmarshal(body, &temp)
	if err != nil {
		fmt.Println("GetTokenErr:", err)
		return ""
	}
	token := temp["access_token"]

	return fmt.Sprintf("%s", token)
}
