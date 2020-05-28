package Marisfrolg_utils

import (
	"encoding/json"
	"fmt"
	"github.com/DeanThompson/jpush-api-go-client"
	"github.com/DeanThompson/jpush-api-go-client/common"
	"github.com/DeanThompson/jpush-api-go-client/device"
	"github.com/DeanThompson/jpush-api-go-client/push"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strings"
	"encoding/base64"
)

type JPushDevice struct {
	Id bson.ObjectId `bson:"_id"`
	//注册设备ID 注册后不会变
	RegistrationId string `bson:"registration_id"`
	//登录者Name
	EmployeeName string `bson:"employee_name"`
	//登录者工号
	EmployeeNo string `bson:"employee_no"`
	////部门ID
	DepartmentId string `bson:"department_id"`
	//标签 可以多个
	Tag []string `bson:"tag"`
	//别名只能一个
	Alias string `bson:"alias"`
	//电话
	Phone string `bson:"phone"`
	//注册设备平台
	RegistrationPlatform string `bson:"registration_platform"`
	//ODS APP内核版本
	OdsAppKernelVersion string `bson:"odsapp_kernel_version"`
	//登录地址
	OdsAppUIVersion string `bson:"odsapp_ui_version"`
	//是否有效
	IsEffective bool `bson:"IsEffective"`
	//创建时间
	CreateTime time.Time `bson:"create_time"`
	//修改时间
	UpdateTime time.Time `bson:"update_time"`
}
type JgVerifyResponse struct {
	Id float64 `json:"id"`
	Code int `json:"code"`
	Content string `json:"content"`
	ExID string `json:"exID"`
	Phone string `json:"phone"`
}

//极光推送相关参数
const (
	AppKey       = "492031ec9c74aa02939b16bc"
	MasterSecret = "7a5d382621ce12a23be1f8e3"
	PushUrl      = "https://bjapi.push.jiguang.cn/v3/push"
	DeviceUrl    = "https://bjapi.push.jiguang.cn/v3/device"
	BASE64_TABLE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	JgLoginTokenVerifyUrl  = "https://api.verification.jpush.cn/v1/web/loginTokenVerify"
	JgVerifyPhoneUrl = "https://api.verification.jpush.cn/v1/web/verify"
    Pubkey       = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDbdXmjcYOtiiUuKipzM4WpIO7t
TQfJvY9Q5U5yYI1MeMba07kg2UG0JNN0190E0/h+FpxF3FNzGqlV+/O7jo7syLI2
E8XBt1lZhGz36bhF/um8QGQ2kwrFHzMWnpmCT0LhMf1A+oeo1De4zmxNmFcZO1gb
8ZbMZiEKACf7+KEBJwIDAQAB
-----END PUBLIC KEY-----
`
    Pirvatekey   = `-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBANt1eaNxg62KJS4q
KnMzhakg7u1NB8m9j1DlTnJgjUx4xtrTuSDZQbQk03TX3QTT+H4WnEXcU3MaqVX7
87uOjuzIsjYTxcG3WVmEbPfpuEX+6bxAZDaTCsUfMxaemYJPQuEx/UD6h6jUN7jO
bE2YVxk7WBvxlsxmIQoAJ/v4oQEnAgMBAAECgYAp5Qg+kmn/1BJ6+KO38EsA2X+j
H4RwF9bnK49JOHNg+OGFXsvFoJPxbuJLOPZBeLHEaE6W65Omsp1HA90onfNcl/vp
qP0a6r6Bf2pDvBBULXqWNnjVpXGHKN4MufEu2bgyhK0O9pg6sI9kCpM/qdWRP5ly
/v2WM0/ryV1URTRzWQJBAPTdXr15MWKKx2TVOMdhn51sLjt3s89ld+x7eX9XtEiF
7WFgMWaWtemUlWT6gXllBZ5FyhFf2ANupzamKxzlUtsCQQDlcFbRPVJQZiC5g0LX
XUuN0OHfDDl2tH4S38S8l1cVLuu8KQH/jAEKJacwFWgUm4qJBmYVUnOeyoyqwBbT
0e6lAkEApEuoVs9rcGgXk7NxTm5VT6YXezU9A6pcheLvSZ9KSuL5vL1zSBdVZa2Z
c9CVcSN0WpcPFwtNADiNn6BtCw1fwwJBAJ0Ov39QGM7Mek5DWjgOty+G83c56QQn
Hb5Ry1zFxGjNy7Tr5WBHOFb323CA1tR0fOq7pJmn7VmfkZc5EudA57kCQHbUrdpM
BwllQWsSIuJnVDwF+l0d8equV9p8oJeweOxWwcLRVyN1FtN5wg6BCYpAzfoR5F5F
ovDYbWPcUnolZh0=
-----END PRIVATE KEY-----
`
)

var client = jpush.NewJPushClient(AppKey, MasterSecret)

//发送推送消息
func PushToDevice(userCode, pushPlatform, logTitle, title, notice string, extraMap map[string]interface{}) (err error) {
	var (
		session *mgo.Session
		c       *mgo.Collection
		list    []*JPushDevice

		ids                 []string
		platform            *push.Platform
		audience            *push.Audience
		notification        *push.Notification
		androidNotification *push.AndroidNotification
		iosNotification     *push.IosNotification
		option              *push.Options
		payload             *push.PushObject
		result              *push.PushResult
		userArr             []string
	)
	if userCode == "" {
		err = fmt.Errorf("请输入要推送的工号，可以是多少个逗号分隔。\n")
		return
	}

	if session, err = mgo.Dial(MONGODB); err != nil {
		goto ERR
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// platform 对象
	platform = push.NewPlatform()
	// 用 Add() 方法添加具体平台参数，可选: "all", "ios", "android"
	//platform.Add("ios")
	// 或者用 All() 方法设置所有平台
	//platform.All()
	if pushPlatform == "all" {
		platform.Add("ios", "android")
	} else {
		platform.Add(pushPlatform)
	}
	// audience 对象，表示消息受众
	audience = push.NewAudience()
	//audience.SetTag([]string{"广州", "深圳"})   // 设置 tag 并集
	//audience.SetTagAnd([]string{"北京", "女"}) // 设置 tag_and 交集
	// audience.SetRegistrationId([]string{"id1", "id2"})   // 设置 registration_id
	// 和 platform 一样，可以调用 All() 方法设置所有受众
	//audience.All()
	//根据用户工号找所有设备ID
	if userCode != "all" {
		userArr = strings.Split(userCode, ",")
		c = session.DB("ODSAPP").C("JPushDevice")
		if err = c.Find(bson.M{"employee_no": bson.M{"$in": userArr}}).All(&list); err != nil {
			goto ERR
		}
		for _, v := range list {
			ids = append(ids, v.RegistrationId)
		}
		audience.SetRegistrationId(ids)
	} else {
		audience.All()
	}

	if len(ids) > 0 || userCode == "all" {
		// notification 对象，表示 通知，传递 alert 属性初始化
		notification = push.NewNotification(logTitle) //推送日志标题

		androidNotification = push.NewAndroidNotification(notice) // android 平台专有的 notification，用 alert 属性初始化
		iosNotification = push.NewIosNotification(notice)         // iOS 平台专有的 notification，用 alert 属性初始化
		// addExtra方法
		for k, v := range extraMap {
			androidNotification.AddExtra(k, v)
			iosNotification.AddExtra(k, v)
		}
		androidNotification.Title = title
		notification.Android = androidNotification

		iosNotification.Badge = 1
		iosNotification.Sound = "default"
		notification.Ios = iosNotification

		option = push.NewOptions()
		option.ApnsProduction = true

		// 可以调用 AddExtra 方法，添加额外信息
		// message.AddExtra("key", 123)
		//创建PushObject对象
		payload = push.NewPushObject()
		payload.Platform = platform
		payload.Audience = audience
		payload.Notification = notification
		//payload.Options = option

		// 打印查看 json 序列化的结果，也就是 POST 请求的 body
		data, err := json.Marshal(payload)
		if err != nil {
			err = fmt.Errorf("json.Marshal PushObject failed:%s\r\n", err)
			goto ERR
		} else {
			fmt.Println("payload:", string(data))
		}
		//开始推送
		result, err = client.Push(payload)
		if err != nil {
			err = fmt.Errorf("Push failed:%s\r\n", err)
			goto ERR
		} else {
			fmt.Println("Push result:", result)
		}
	}

	return
ERR:
	return
}

//设置标签
func SetAliaxByDevice(c *mgo.Collection, obj *JPushDevice) (err error) {
	var (
		result       *common.ResponseBase
		list         []*JPushDevice
		deviceUpdate *device.DeviceUpdate
	)

	if err = c.Find(bson.M{"employee_no": obj.EmployeeNo}).All(&list); err != nil {
		return
	}
	if list != nil {
		deviceUpdate = device.NewDeviceUpdate()
		for _, v := range list {
			if v.Alias == "" || v.Alias != v.EmployeeNo {
				deviceUpdate.SetAlias(obj.EmployeeNo)
				deviceUpdate.SetMobile(obj.Phone)
				result, err = client.UpdateDevice(v.RegistrationId, deviceUpdate)
				showResultOrError("client.UpdateDevice", result, err)
			}

		}
	}

	return
}

func showResultOrError(method string, result interface{}, err error) {
	if err != nil {
		fmt.Printf("%s failed: %v\r\n", method, err)
	} else {
		fmt.Printf("%s result: %v\r\n", method, result)
	}
}

func JgBase64EncodeToString()(str string){
	input := []byte(fmt.Sprintf(`%s:%s`,AppKey,MasterSecret))

	// 演示base64编码
	str= base64.StdEncoding.EncodeToString(input)
	fmt.Println(str)
	return str
}