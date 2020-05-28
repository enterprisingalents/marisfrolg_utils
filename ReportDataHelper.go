package Marisfrolg_utils

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"gopkg.in/mgo.v2"
)

type _BrandWithCityAll struct {
	//第二部分 城市&渠道分组 结构体
	WERKS          string
	KTOKD          string
	SHOPNAME       string
	BEHVO          string
	ZoneName       string
	LargeZone      string
	CityName       string
	CC_YEAR_SUM    float64 //当年（金额）
	CC_MONTH_SUM   float64 //当月（金额）
	CC_DAY_SUM     float64 //当日（金额）
	CC_TBYEAR_SUM  float64 //上年（金额）
	CC_TBMONTH_SUM float64 //同比月（金额）
	CC_HBMONTH_SUM float64 //环比月（金额）

	CC_YEAR_SUM_Q    float64 //当年（数量）
	CC_MONTH_SUM_Q   float64 //当月（数量）
	CC_DAY_SUM_Q     float64 //当日（数量）
	CC_TBYEAR_SUM_Q  float64 //上年（数量）
	CC_TBMONTH_SUM_Q float64 //同比月（数量）
	CC_HBMONTH_SUM_Q float64 //环比月（数量）
	CC_RANGE_SUM     float64 //范围值
	CC_RANGE_SUM_Q   float64 //范围数量值
}


//只关联了用到的地方，差的内容请去Account.go去复制过来
type Loginer struct {
	Id                bson.ObjectId `bson:"_id"`
	EmployeeNo        string        `bson:"ID"`
	Password          string        `bson:"password"`
	Name              string        `bson:"eName"`
	IsSetPwd          string        `bson:"isSetPwd"`      //1同意公司协议
	WorkState         string        `bson:"eWorkState"`    //1在职，2离职
	SonWorkState      string        `bson:"eSonWorkState"` //1正常，2异常，3停薪离职
	PhotoUrl          string
	RequstTime        time.Time
	AccessToken       string         //安全令牌
	MockCode          string         //模拟登陆用的唯一码
	MyPermission      []string       //我的权限
	MyCompany         []string       //我的公司范围
	MyBrand           []string       //我的品牌范围
	MyLargeArea       []string       //我的大区范围
	MyShop            []ShopOrOffice //我的物理门店
	MyOffice          []ShopOrOffice //我的物理办事处
	MyPublisher       []string       //我的发布方范围
	MyUserOrg         []string       //人事系统权限
	IsAgree           bool           //是否同意软件协议
	DepId             string //部门ID
	Tag               string //部门名称
}

type ShopOrOffice struct {
	Code string
	Name string
}

type UserBrand struct {
	Id         bson.ObjectId `bson:"_id"`
	EmployeeNo string        `bson:"employeeNo"`
	BrandCode  string        `bson:"brandCode"`
}

type UserLargeAndShop struct {
	Id           bson.ObjectId       `bson:"_id"`
	EmployeeNo   string              `bson:"employeeNo"`
	LargeAndShop []*ShowLargeAndShop `bson:"largeAndShop"`
}

type ShowLargeAndShop struct {
	Code       string
	Name       string
	Checked    bool //界面对应勾选
	AllChecked bool //是否全选
	Selected   bool //是否拥有权限
	SalesBrand string
	Status     int // 0 原有  1新增 2 删除
	Children   []*ShowLargeAndShop
}

//根据门店权限筛选数据
func FilterDataByLoginerOfLargeAndShop(yourNumber string, orilist []_BrandWithCityAll) (list []_BrandWithCityAll, err error) {
	//获取所有权限
	loginer, err := GetAreaAndShopRange(yourNumber)
	oldLargeList := make(map[string]ShopOrOffice) //大区旧值字典
	for _, v := range loginer.MyShop {
		oldLargeList[v.Code] = v
	}

	for _, row := range orilist {
		if _, has := oldLargeList[row.WERKS]; has {
			list = append(list, row)
		}
	}

	return
}

//获取大区门店范围
func GetAreaAndShopRange(empNo string) (loginer *Loginer, err error) {
	session, err := mgo.Dial(MONGODB)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	loginer = new(Loginer)

	//获取品牌列表
	c := session.DB("ODSAPP").C("empNo")
	queryOldBrand := c.Find(bson.M{"employeeNo": empNo})
	var oldBrandList []*UserBrand
	queryOldBrand.All(&oldBrandList)
	for _, v := range oldBrandList {
		loginer.MyBrand = append(loginer.MyBrand, v.BrandCode)
	}

	//获取大区与物理门店
	c = session.DB("ODSAPP").C("UserLargeAndShop")
	queryLargeAndShop := c.Find(bson.M{"employeeNo": empNo})
	var oldLargeAndShop UserLargeAndShop
	queryLargeAndShop.Limit(1).One(&oldLargeAndShop)

	for _, v := range oldLargeAndShop.LargeAndShop {
		if v.Selected {
			loginer.MyLargeArea = append(loginer.MyLargeArea, v.Code)
		}
		for _, w := range v.Children {
			if w.Selected {
				tem := new(ShopOrOffice)
				tem.Code = w.Code
				tem.Name = w.Name
				if w.Name != "已删除门店" {
					loginer.MyShop = append(loginer.MyShop, *tem)
				}
			}
		}
	}

	return
}
