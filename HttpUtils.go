package Marisfrolg_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Text struct {
	Content string `json:"content"`
}

type Textcard struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
type File struct {
	Media_id string `json:"media_id"`
}
type Message struct { //touser、toparty、totag不能同时为空
	Touser   string   `json:"touser"` //员工工号，多人用|隔开。(如：00275|01953)
	Toparty  string   `json:"toparty"`
	Totag    string   `json:"totag"`
	Msgtype  string   `json:"msgtype"`
	Agentid  int      `json:"agentid"`  //企业应用的id
	Safe     int      `json:"safe"`     //表示是否是保密消息
	Text     Text     `json:"text"`     //Msgtype为text时方可使用
	File     File     `json:"file"`     //Msgtype为file时方可使用
	Textcard Textcard `json:"textcard"` //Msgtype为textcard时方可使用
}
type UploadMediaRes struct {
	Errcode  int
	Errmsg   string
	Type     string
	Media_id string
}

// get 网络请求
func HttpGet(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlstr := Url.String()
	resp, err := http.Get(urlstr)
	//fmt.Println(urlstr)
	if err != nil {
		//fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// post 网络请求 ,params 是url.Values类型
func HttpPost(apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func HttpPost2(apiURL string, params url.Values, parmbody string, token string) (rs []byte, err error) {

	//req, _ := http.NewRequest("POST", apiURL, nil)
	//strings.NewReader("name=cjb")  "templateCode =车辆管理&businessObject ={}&"
	//req, err := http.NewRequest("POST", apiURL, strings.NewReader("templateCode=车辆管理&businessObject={}&UserCode=02966&UserName=02999"))

	data := params.Encode()

	req, err := http.NewRequest("POST", apiURL+"?"+data, strings.NewReader(parmbody))

	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Content-Type", "application/json")
	// cer := url.Values{}
	// cer.Set("Node ", parmbody)
	// req.PostForm = cer

	//hc := http.Client{}
	res, err := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(body))

	//AccessToken, _ := getBearer("F7894ADD5867E4DF")
	// resp, err := http.PostForm(apiURL, params)
	// resp.Header.Add("authorization", "Bearer "+token)
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()
	return body, err

}

// get 网络请求
func HttpGet2(apiURL string, params url.Values, token string) (rs []byte, err error) {
	data := params.Encode()

	req, err := http.NewRequest("GET", apiURL+"?"+data, nil)
	req.Header.Add("authorization", "Bearer "+token)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	//fmt.Println(string(body))
	return body, err

}

//http://192.168.2.14/CompanyWXPlat/QyWeiXin/SendMessage?ValidationCode=ZYAEGPXTNPI7RLPM&Agentid=0&UserID=00275&Content=123
func SendMessage(userCode string, Message string) (err error, result string) {
	//请求地址
	juheURL := "http://192.168.2.14/CompanyWXPlat/QyWeiXin/SendMessage?"
	//初始化参数
	param := url.Values{}
	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	param.Set("ValidationCode", "ZYAEGPXTNPI7RLPM")
	param.Set("Agentid", "0")
	param.Set("UserID", userCode)
	param.Set("Content", Message)

	//发送请求
	data, err := HttpGet(juheURL, param)
	if err != nil {
		fmt.Printf("请求失败,我方错误信息:\r\n%v", err)
	} else {
		result = string(data)
	}

	return
}

/*
测试文本
{
	"agentid": 0,
	"msgtype": "text",
	"safe": 0,
	"text": {
	  "content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。"
	},
	"touser": "00275"
  }
测试卡片
{
	"touser" : "00275|01953",
	"msgtype" : "textcard",
	"agentid" : 22,
	"textcard" : {
			 "title" : "领奖通知",
			 "description" : "<div class=\"gray\">2016年9月26日</div> <div class=\"normal\">恭喜你抽中iPhone 7一台，领奖码：xxxx</div><div class=\"highlight\">请于2016年10月10日前联系行政同事领取</div>",
			 "url" : "URL"
	},
	"safe":1
 }
测试文件
{
  "agentid": 22,
  "msgtype": "file",
  "safe": 0,
   "file" : {
		"media_id" : "3uddU0MrKnqxy1hKXA9O9oBIUJ4n7b1UgWbby-jGqfo34OuJ5Y_phEuucENi-3Pv1"
   },
  "touser": "00275"
}

*/

/*
带文件的消息推送
path:网络地址(如：)
usercode：工号(如：02607)
*/
func PushMessageWithFile(path, token, usercode string) (err error) {
	var data []byte
	var UM *UploadMediaRes
	var F File
	var mege Message
	LocalhostUrl := "http://192.168.2.121:8998/v1/Push/UploadMedia" //正式
	//LocalhostUrl := "http://localhost:8998/v1/Push/UploadMedia" //本地测试用
	param := url.Values{}

	param.Set("fileUrl", path)
	if data, err = HttpPost2(LocalhostUrl, param, "", token); err != nil {
		err = fmt.Errorf("请求失败,上传临时素材出错；返回错误信息=%s,data=%s", err, string(data))
		return err
	}
	if err = json.Unmarshal(data, &UM); err != nil {
		err = fmt.Errorf("获取临时素材id转json时出错;json.Unmarshal错误信息=%s,data=%s", err, string(data))
		return
	}

	F.Media_id = UM.Media_id

	mege.Msgtype = "file"
	mege.Touser = usercode
	mege.Agentid = 0
	mege.Safe = 0
	mege.File = F

	_, err = PushMessage(mege, token)
	if err != nil {
		x := fmt.Sprintf("%s", err)
		err = errors.New("推送消息时报错；错误消息为：" + x + "。")
	}
	return
}

//go语言发企业微信消息
func PushMessage(meg Message, token string) (result []byte, err error) {
	apiURL := "http://192.168.2.121:8998/v1/Push/PushMessage"
	jdata, err := json.Marshal(meg)
	data := string(jdata)
	bearer := "Bearer " + token
	result, err = HttpPostFuc(apiURL, bearer, data, "application/json")
	if err != nil {
		fmt.Printf("PushMessage HttpPostFuc err:%s",err)
		return
	}

	return
}

//go语言发企业微信消息
func PushMessageByJsonString(msg string, token string) (result []byte, err error) {
	//params := url.Values{}
	apiURL := "http://192.168.2.121:8998/v1/Push/PushMessage"
	bearer := "Bearer " + token
	result, err = HttpPostFuc(apiURL, bearer, msg, "application/json")
	if err != nil {
		return
	}

	return
}
func HttpPostFuc(url, token, postData, contentType string) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(postData))
	if token!=""{
		request.Header.Set("Authorization", token)
	}
	request.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("http.Get err %s", err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求没有成功，请求状态为:%d\n", resp.StatusCode)
	}

	if bytes, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll err:%s\n", err.Error())
	} else {
		return bytes, nil
	}

}


