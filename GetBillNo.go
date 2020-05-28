package Marisfrolg_utils

import (
	"database/sql"
	"fmt"
	"strings"
	//"reflect"
	//"strings"
	"strconv"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

type UserPermissionRequest struct {
	BillNo string `bson:"billNo"`
}

type LookBookRequest struct {
	BillNo string `bson:"billNo"`
}

var (
	//获取权限订单编码
	bill_qx    = 0                             //权限申请单号初始化
	bill_qxymd = time.Now().Format("20060102") //权限申请单号年月日
	//获取指导手册订单编码
	bill_lbq    = 0
	bill_lbqymd = time.Now().Format("20060102")

	//直播单号
	bill_zb    = 0
	bill_zbymd = time.Now().Format("20060102")
)

func GetQXBillNo() (no string, err error) {
	var mutex sync.Mutex
	var No string
	ymd := time.Now().Format("20060102")
	defer func() {
		bill_qxymd = ymd
		fmt.Println("Unlock get Id:", bill_qx)
		mutex.Unlock()
	}()
	mutex.Lock()

	//初始化
	if bill_qx == 0 {
		session, err := mgo.Dial(MONGODB)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		var request *UserPermissionRequest
		var BillNo string
		c1 := session.DB("ODSAPP").C("UserPermissionRequest")
		queryOld := c1.Find(nil).Sort("-billNo").Skip(0).Limit(1)
		err = queryOld.One(&request)
		if err == nil {
			BillNo = request.BillNo
			//已存在单号 +1
			if BillNo != "" {
				n := SubString(BillNo, len(BillNo)-5, len(BillNo))
				d := SubString(BillNo, 2, 8)
				if strings.Compare(bill_qxymd, ymd) == 0 { //同一天
					if strings.Compare(ymd, d) == 0 {
						id, _ := strconv.Atoi(n)
						bill_qx = bill_qx + 1
						bill_qx = bill_qx + id
						fmt.Println("Unlock get Id:", bill_qx)
					} else {
						bill_qx = 1
						fmt.Println("Unlock get Id:", bill_qx)
					}
				} else {
					bill_qx = 1
					fmt.Println("Unlock get Id:", bill_qx)
				}

			} else {
				bill_qx = bill_qx + 1
				fmt.Println("Unlock get Id:", bill_qx)
			}
		}
	} else {
		if strings.Compare(bill_qxymd, ymd) == 0 {
			bill_qx = bill_qx + 1
			fmt.Println("Unlock get Id:", bill_qx)
		} else { //新一天重置 r
			bill_qx = 1
			fmt.Println("Unlock get Id:", bill_qx)
		}
	}
	No = fmt.Sprintf("QX%s%s", ymd, PadLeft(strconv.Itoa(bill_qx), 5, "0"))
	return No, err
}

//获取指导手册订单编码
func GetLookBookRequestBillNo() (no string, err error) {
	var (
		mutex   sync.Mutex
		session *mgo.Session
		c       *mgo.Collection
	)
	ymd := time.Now().Format("20060102")
	request := new(LookBookRequest)
	billNo := ""
	defer func() {
		bill_lbqymd = ymd
		fmt.Println("Unlock 获取最新指导手册订单号:", no)
		mutex.Unlock()
	}()
	mutex.Lock()

	//初始化
	if bill_lbq == 0 {
		session, err = mgo.Dial(MONGODB)
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c = session.DB("ODSAPP").C("LookBookRequest")
		c.Find(nil).Sort("-billNo").Limit(1).One(&request)
		if request.BillNo != "" {
			billNo = request.BillNo
			//已存在单号 +1
			if billNo != "" {
				n := SubString(billNo, len(billNo)-5, len(billNo))
				d := SubString(billNo, 2, 8)
				if strings.Compare(bill_lbqymd, ymd) == 0 { //同一天
					if strings.Compare(ymd, d) == 0 {
						id, _ := strconv.Atoi(n)
						bill_lbq = bill_lbq + 1
						bill_lbq = bill_lbq + id
						fmt.Println("Unlock get LQB Id:", bill_lbq)
					} else {
						bill_lbq = 1
						fmt.Println("Unlock get LQB Id:", bill_lbq)
					}
				} else {
					bill_lbq = 1
					fmt.Println("Unlock get LQB Id:", bill_lbq)
				}

			} else {
				bill_lbq = bill_lbq + 1
				fmt.Println("Unlock get LQB Id:", bill_lbq)
			}
		}
	} else {
		if strings.Compare(bill_lbqymd, ymd) == 0 {
			bill_lbq = bill_lbq + 1
			fmt.Println("Unlock get LQB Id:", bill_lbq)
		} else { //新一天重置 r
			bill_lbq = 1
			fmt.Println("Unlock get LQB Id:", bill_lbq)
		}
	}
	no = fmt.Sprintf("LQ%s%s", ymd, PadLeft(strconv.Itoa(bill_lbq), 5, "0"))
	return
}

//直播单号
func GetSkuOrderBillNo() (no string, err error) {
	var (
		mutex sync.Mutex
	)
	ymd := time.Now().Format("20060102")
	//var billNo string
	defer func() {
		bill_zbymd = ymd
		//fmt.Println("Unlock 直播单号:", no)
		mutex.Unlock()
	}()
	mutex.Lock()

	//初始化
	if bill_zb == 0 {
		db, err := sql.Open("mysql", ORDER_ONLINE)
		defer db.Close()
		var billNo string
		sql := fmt.Sprintf(`SELECT IFNULL(max(billno),'') FROM orderlist`)
		db.QueryRow(sql).Scan(&billNo)

		if err == nil {
			//已存在单号 +1
			if billNo != "" {
				n := SubString(billNo, len(billNo)-5, len(billNo))
				d := SubString(billNo, 2, 10)
				if strings.Compare(bill_zbymd, ymd) == 0 { //同一天
					if strings.Compare(ymd, d) == 0 {
						id, _ := strconv.Atoi(n)
						bill_zb = bill_zb + 1
						bill_zb = bill_zb + id
						fmt.Println("Unlock get LQB Id:", bill_zb)
					} else {
						bill_zb = 1
						fmt.Println("Unlock get LQB Id:", bill_zb)
					}
				} else {
					bill_zb = 1
					fmt.Println("Unlock get LQB Id:", bill_zb)
				}

			} else {
				bill_zb = bill_zb + 1
				fmt.Println("Unlock get LQB Id:", bill_zb)
			}
		}
	} else {
		if strings.Compare(bill_zbymd, ymd) == 0 {
			bill_zb = bill_zb + 1
			fmt.Println("Unlock get LQB Id:", bill_zb)
		} else { //新一天重置 r
			bill_zb = 1
			fmt.Println("Unlock get LQB Id:", bill_zb)
		}
	}
	no = fmt.Sprintf("ZB%s%s", ymd, PadLeft(strconv.Itoa(bill_zb), 5, "0"))
	return no, err
}
