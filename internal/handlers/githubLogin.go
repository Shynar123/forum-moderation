package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"forum/internal/cookies"
	"forum/internal/types"
)

const (
	githubAuthURL     = "https://github.com/login/oauth/authorize"
	githubTokenURL    = "https://github.com/login/oauth/access_token"
	githubUserInfoURL = "https://api.github.com/user"
)

var configGithub = &types.Config{
	ClientID:     "cf963198dc34211c3405",
	ClientSecret: "3b38d5fb44e1614babb6b1b644ff7ca7a17a57a8",
	Endpoint:     "http://localhost:8000/github/callback",
	Scopes:       "email profile",
}

type githubInfo struct {
	Username string `json:"login"`
	pass     string `json:"node_id"`
}

func (h *Handler) handleGithubLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}
	url := fmt.Sprintf("%s?scope=user:email&client_id=%s", githubAuthURL, configGithub.ClientID)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *Handler) handleGithubCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	data := strings.NewReader(fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&accept=:json", code, configGithub.ClientID, configGithub.ClientSecret))
	resp, err := http.Post(githubTokenURL, "application/x-www-form-urlencoded", data)
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

	token := h.GetTokenGit(body)


	req, err := http.NewRequest("GET", githubUserInfoURL, nil)
	if err != nil {
		fmt.Println("Google get error:", err)
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

	var githubUser githubInfo
	err = json.Unmarshal(info, &githubUser)
	if err != nil {
		fmt.Println("Unmarshal error ", err)
		return
	}
	user := &types.CreateUserData{
		Username: strings.ReplaceAll(githubUser.Username, " ", ""),
		Password: githubUser.pass,
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

func (h *Handler) GetTokenGit(body []byte) string {
	params, err := url.ParseQuery(string(body))
	if err != nil {
		return ""
	}

	accessToken := params.Get("access_token")
	return accessToken
}
