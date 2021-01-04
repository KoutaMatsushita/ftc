package main

import (
	"bytes"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type FirebaseTokenRequest struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type FirebaseTokenResponse struct {
	IdToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

var app *firebase.App
var client *auth.Client

func main() {
	ctx := context.Background()
	var (
		uid = flag.String("uid", "uid", "レスポンスの jwt の uid")
		apikey = flag.String("apikey", "", "認証に使う firebase プロジェクトの apikey")
	)
	flag.Parse()
	if *apikey == "" {
		_apikey := os.Getenv("FTC_APIKEY")
		apikey = &_apikey
	}

	if *apikey == "" {
		msg :=
			`apiKey が見つかりません。
環境変数 FTC_APIKEY をセットするか、引数で渡してください。
詳しくは https://firebase.google.com/docs/reference/rest/auth/ を参照してください。`
		fmt.Println(msg)
		os.Exit(1)
	}

	initApp(ctx)
	initClient(ctx)

	customToken := getCustomToken(ctx, uid)
	token := getToken(apikey, customToken)

	fmt.Println(token.IdToken)
}

// firebase の初期化
func initApp(ctx context.Context) {
	a, err := firebase.NewApp(ctx, nil)
	if err != nil {
		panic(err)
	}
	app = a
}

// firebase 認証クライアントの初期化
func initClient(ctx context.Context) {
	c, err := app.Auth(ctx)
	if err != nil {
		panic(err)
	}
	client = c
}

// firebase からカスタムトークンを取得
func getCustomToken(ctx context.Context, uid *string) string {
	token, err := client.CustomToken(ctx, *uid)
	if err != nil {
		panic(err)
	}

	return token
}

// firebase からアクセストークンを取得
func getToken(apikey *string, token string) *FirebaseTokenResponse {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%v", *apikey)
	requestBody, err := json.Marshal(FirebaseTokenRequest{Token: token, ReturnSecureToken: true})
	if err != nil {
		panic(err)
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	if response.StatusCode >= 400 {
		panic(fmt.Errorf(string(b)))
	}

	ft := FirebaseTokenResponse{}
	err = json.Unmarshal(b, &ft)
	if err != nil {
		panic(err)
	}

	return &ft
}
