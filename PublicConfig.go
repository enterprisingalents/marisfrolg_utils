package Marisfrolg_utils

import (
	"strconv"

	"github.com/larspensjo/config"
)

var CACHE_CONN = "root:grace1996erp@tcp(192.168.2.119:3306)/ods_cache?parseTime=true"
var ORDER_ONLINE = "root:grace1996erp@tcp(192.168.2.119:3306)/orderonline?parseTime=true"
var MYSQL_DB_CONNECT = "train:train123@tcp(192.168.2.118:3306)/ods_train?parseTime=true"
var PLAN_CONNECT = "plan:plan123@tcp(192.168.2.118:3306)/ods_sales_plan?parseTime=true"
var MONGODB = "mongodb://root:Grace1996db@192.168.2.85:27017/"
var ORACLE_CONN_CONNECT = "mfdev/grace1996erp@M4DEV"
var REDIS_CONN = "192.168.11.234:6379"
var WF_CONN = "http://192.168.11.234:8999/v1/WorkFlow"
var HR_CONN = "http://192.168.11.234:7090/v1"
var WXPlatform_CONN = "http://192.168.11.234:7996/v1/UserInfo"
var PORT = 8766
var InvitationNoticeUserList = "02730"
var QRCODE_IP="192.168.11.234"
var QRCODE_PORT=7997

const (
	HANA_DRIVER = "hdb"
	// HANA_DNS    = "hdb://LUSHUAIHUA:llsshhLSH123@210.75.9.164:30015"
	HANA_DNS = "hdb://SYSTEM:Grace12345@192.168.2.34:30015"
)

const (
	ORA_DRIVER  = "oci8"
	ORA_URIbank = "system/information@MUTIBANK"
	//ods_bak/odsbak@110.1.5.54:1521/ODS_DEV
)

//销售总表相关参数
const (
	ODSReportGetTokenUrl              = "https://shopservice.sydj520.com/Authorize/Token"
	ODSReportPostUrl                  = "https://shopservice.sydj520.com/udx/service"
	ODSReportClientId                 = "ODS"
	ODSReportClientSecret             = "ODSServiceLive"
	ODSReportUniqueKey                = "FitShop.FitODSReportDomain"
	ODSReportSalesAndRefundMethodName = "GetSalesAndRefund"
	ODSReportRefundMethodName         = "GetRefund"
	ODSReportSalesMethodName          = "GetSales"
)

//fir.im参数
const (
	MA              = "5df84a82f94548378b6698b4"
	SU              = "5df84a82f94548378b6698b4"
	AUM             = "5df84a82f94548378b6698b4"
	ZH              = "5df84a82f94548378b6698b4"
	FirUrl          = "http://api.fir.im/apps/latest/"
	ApiToken        = "51fa892530359ecb0f3bec06f3a83f01"
	ODSAndroidAppId = "5dcd0fb3f94548665e9f0a22"
	ODSIOSAppId     = "5dcd0f4923389f6f8018a696"
)

var CurrentMode = "DEV"

func RunMode(Mode string) {
	conf, _ := config.ReadDefault("conf/app.conf")
	CACHE_CONN, _ = conf.String(Mode, "CACHE_CONN")
	MYSQL_DB_CONNECT, _ = conf.String(Mode, "MYSQL_CONN")
	ORACLE_CONN_CONNECT, _ = conf.String(Mode, "ORACLE_CONN")
	PLAN_CONNECT, _ = conf.String(Mode, "PLAN_CONN")
	MONGODB, _ = conf.String(Mode, "MONGO_CONN")
	REDIS_CONN, _ = conf.String(Mode, "REDIS_CONN")
	WF_CONN, _ = conf.String(Mode, "WF_CONN")
	HR_CONN, _ = conf.String(Mode, "HR_CONN")
	WXPlatform_CONN, _ = conf.String(Mode, "WXPlatform_CONN")
	_PORT, _ := conf.String(Mode, "PORT")
	PORT, _ = strconv.Atoi(_PORT)
	InvitationNoticeUserList, _ = conf.String("Common", "InvitationNoticeUserList")
	ORDER_ONLINE, _ = conf.String(Mode, "ORDER_ONLINE")
	switch CurrentMode {
	case "DEV":
		QRCODE_IP = "192.168.11.234"
		QRCODE_PORT = 7997
	case "TEST":
		QRCODE_IP = "192.168.11.234"
		QRCODE_PORT = 7997
	case "PRD":
		QRCODE_IP = "192.168.2.121"
		QRCODE_PORT = 8997
	}
}
