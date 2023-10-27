package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"net/http"
	"net/url"
	"time"
)

type WechatClient struct {
	httpClient *http.Client
}

func NewWechatClient(httpClient *http.Client) *WechatClient {

	if httpClient == nil {

		httpClient = http.DefaultClient
		httpClient.Timeout = time.Second * 5
	}

	return &WechatClient{
		httpClient: httpClient,
	}
}

func (pc *WechatClient) WXLogin(jscode string) (wxMap map[string]string, error error) {

	loginUrl := fmt.Sprintf(global.GVA_CONFIG.Wechat.SessionUrl,
		url.QueryEscape(global.GVA_CONFIG.Wechat.AppID), url.QueryEscape(global.GVA_CONFIG.Wechat.Secret), url.QueryEscape(jscode))
	fmt.Println("----loginUrl:", loginUrl)
	httpResp, err := pc.httpClient.Get(loginUrl)

	if err != nil {
		return wxMap, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return wxMap, fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	err = json.NewDecoder(httpResp.Body).Decode(&wxMap)
	fmt.Println("----wxMap:", wxMap)

	if err != nil {
		fmt.Println("----err:", err)
	}

	//err = json.NewDecoder(httpResp.Body).Decode(&session)
	////respData := NewPayData()
	////err = respData.FromJson(httpResp.Body)
	//if err != nil {
	//	fmt.Println("--1--session:", session)
	//	fmt.Println("--1--err:", err)
	//	return session, err
	//}

	return wxMap, nil
}
