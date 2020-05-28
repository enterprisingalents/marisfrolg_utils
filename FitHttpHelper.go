package Marisfrolg_utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type TokenResponse struct {
	TokenType   string        `json:"token_type"`
	ClientId    string        `json:"client_id"`
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Duration `json:"expires_in"`
	ShopGid     string        `json:"shopGid"`
	ShopName    string        `json:"shopName"`
	UserGid     string        `json:"UserGid"`
}

func GetShopAuthorizeToken(GetTokenUrl, client_id, client_secret string) (resObj *TokenResponse, err error) {
	var (
		reqUrl   string
		response *http.Response
		bytes    []byte
		tokenObj TokenResponse
	)
	reqUrl = fmt.Sprintf(`%s?client_id=%s&client_secret=%s`, GetTokenUrl, client_id, client_secret)
	response, err = http.Get(reqUrl)
	defer response.Body.Close()
	if err != nil {
		err = fmt.Errorf("http.Get error:%s", err)
		goto ERR
	}
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("error response code:%d\n", response.StatusCode)
		goto ERR
	}
	bytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll error:%s\n", err)
		goto ERR
	}
	//tokenInfo = make(map[string]interface{})
	err = json.Unmarshal(bytes, &tokenObj)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal error:%s\n", err)
	}
	return &tokenObj, err
ERR:
	return nil, err
}

func HttpPostShopFuc(url, token, postData, contentType string) (resBytes []byte, err error) {
	var (
		request *http.Request
		resp    *http.Response
	)

	request, _ = http.NewRequest(http.MethodPost, url, strings.NewReader(postData))
	request.Header.Set("Authorization", "Bearer"+" "+token)
	request.Header.Set("Content-Type", contentType)
	if resp, err = http.DefaultClient.Do(request); err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("请求状态为:%d\n", resp.StatusCode)
		return
	}
	if resBytes, err = ioutil.ReadAll(resp.Body); err != nil {
		err = fmt.Errorf("ioutil.ReadAll err:%s\n", err)
	}
	return
}
