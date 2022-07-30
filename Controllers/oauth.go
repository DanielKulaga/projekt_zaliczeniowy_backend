package Controllers

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"myapp/Models"
	"net/http"
	"os"
)

func generateAccessToken() string {
	token := uuid.New()
	return token.String()
}

func getGoogleConf() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_IDENTIFICATOR_GOOGLE"),
		ClientSecret: os.Getenv("CLIENT_SECRET_GOOGLE"),
		RedirectURL:  "http://localhost:1323/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return conf
}

func getGithubConf() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_IDENTIFICATOR_GITHUB"),
		ClientSecret: os.Getenv("CLIENT_SECRET_GITHUB"),
		RedirectURL:  "http://localhost:1323/github/callback",
		Scopes: []string{
			"user:email",
			"read:user",
		},
		Endpoint: github.Endpoint,
	}

	return conf
}

func GoogleLogin(c echo.Context) error {
	url := getGoogleConf().AuthCodeURL("state")
	return c.JSON(http.StatusOK, url)
}

func GithubLogin(c echo.Context) error {
	url := getGithubConf().AuthCodeURL("state")
	return c.JSON(http.StatusOK, url)
}

func GoogleCallback(c echo.Context) error {

	token, err := getGoogleConf().Exchange(context.Background(), c.QueryParam("code"))

	if err != nil {
		print(err)
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		print(err)
	}
	defer response.Body.Close()

	userInfo, err := ioutil.ReadAll(response.Body)

	if err != nil {
		print(err)
	}

	oauthUser := struct {
		Name  string
		Email string
	}{}
	err = json.Unmarshal(userInfo, &oauthUser)

	if err != nil {
		print(err)
	}

	userToken := generateAccessToken()

	if !FindUserInDB(oauthUser.Email) {
		newUser := Models.User{}

		newUser.Email = oauthUser.Email
		newUser.Username = oauthUser.Name
		newUser.Oauthtoken = token.AccessToken
		newUser.Usertoken = userToken

		AddUser(newUser)
	} else {
		EditUserToken(oauthUser.Email, userToken)
	}

	log.Printf("Redirect")

	c.Redirect(http.StatusFound, "http://localhost:3000/login/success?token="+userToken+"&email="+oauthUser.Email)

	return c.JSON(http.StatusOK, "Login successfully")
}

func GithubCallback(c echo.Context) error {

	token, err := getGithubConf().Exchange(context.Background(), c.QueryParam("code"))

	if err != nil {
		print(err)
	}

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+token.AccessToken)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		print(err)
	}

	defer response.Body.Close()

	userInfo, err := ioutil.ReadAll(response.Body)
	userInfoString := string(userInfo)

	if err != nil {
		print(err)
	}

	print(userInfoString)

	oauthUser := struct {
		ID    int
		Login string
		Email string
	}{}
	err = json.Unmarshal([]byte(userInfoString), &oauthUser)

	if err != nil {
		print(err)
	}

	userToken := generateAccessToken()

	if !FindUserInDB(oauthUser.Email) {
		newUser := Models.User{}
		newUser.Email = oauthUser.Email
		newUser.Username = oauthUser.Login
		newUser.Oauthtoken = token.AccessToken
		newUser.Usertoken = userToken

		AddUser(newUser)
	} else {
		EditUserToken(oauthUser.Email, userToken)
	}

	log.Printf("Redirect")

	c.Redirect(http.StatusFound, "http://localhost:3000/login/success?token="+userToken+"&email="+oauthUser.Email)

	return c.JSON(http.StatusOK, "Login successfully")
}
