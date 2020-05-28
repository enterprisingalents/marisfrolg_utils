package Marisfrolg_utils

import (
	"fmt"

	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

/// <summary>
/// 阿里短信发送接口
/// 短信服务SDK简介 https://help.aliyun.com/document_detail/101874.html?spm=a2c4g.11186623.6.592.78ba5f30ZXZjpI
/// </summary>
/// <param name="PhoneNumbers">
/// 短信接收号码,支持以逗号分隔的形式进行批量调用，
/// 批量上限为800个手机号码,批量调用相对于单条调用及时性稍有延迟,
/// 验证码类型的短信推荐使用单条调用的方式</param>
/// <param name="SignName">短信签名</param>
/// <param name="TemplateCode">短信模板ID</param>
/// <param name="TemplateParam">
/// 可选参数,
/// 短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,
/// 请参照标准的JSON协议对换行符的要求,比如短信内容中包含\r\n的情况在JSON中需要表示成\r\n,
/// 否则会导致JSON在服务端解析失败
/// </param>
/// <returns></returns>
func SendSms(PhoneNumbers string, SignName string, TemplateCode string, TemplateParam string) string {
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", "LTAID6oGoC199rbD", "SOlGbE4MTOmLgyQAASlZ8RGarDxFSo")
	if err != nil {
		panic(err)
	}

	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["PhoneNumbers"] = PhoneNumbers
	request.QueryParams["SignName"] = SignName
	request.QueryParams["TemplateCode"] = TemplateCode
	request.QueryParams["TemplateParam"] = TemplateParam

	response, err := client.ProcessCommonRequest(request)
	if err != nil {
		panic(err)
	}
	fmt.Print(response.GetHttpContentString())

	return response.GetHttpContentString()
}

//使用企业微信发送通知  企业小助手：0
func SendSmsFromEnterpriseWechat(agentid, employeeNos, content string) {
	userCodeArr := strings.Split(employeeNos, ",")
	for _, userId := range userCodeArr {
		requestUrl := "http://192.168.2.14/CompanyWXPlat/QyWeiXin/SendMessage?"
		urlInfo, err := url.Parse(requestUrl)
		params := url.Values{}
		params.Set("ValidationCode", "ZYAEGPXTNPI7RLPM")
		params.Set("Agentid", agentid) //0 企业助手
		params.Set("UserID", userId)
		params.Set("Content", content)
		urlInfo.RawQuery = params.Encode()
		fmt.Println("postData=" + urlInfo.String())
		_, err = http.Get(urlInfo.String())
		if err != nil {
			log.Printf("SendMessage出错:%s \n", err)
		}
	}
}

// 发送通知给QRCODE changedType:PERMISSION_CHANGED MENU_CHANGED BADGE_CHANGED
func SendMessageToQrCode(userId, changedType string) {
	im, _ := gosocketio.Dial(
		gosocketio.GetUrl(QRCODE_IP, QRCODE_PORT, false),
		transport.GetDefaultWebsocketTransport())
	im.Emit(changedType, userId)

}
