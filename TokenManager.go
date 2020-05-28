package Marisfrolg_utils

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/store"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	OAuth2Models "gopkg.in/oauth2.v3/models"
)

var TokenManager *manage.Manager

func GenerateToken(clientID string, clientSecret string) (ti oauth2.TokenInfo, err error) {

	var theUser *TokenUser
	//判断是否是惠州的工号（HZ）
	var IsHZ = SubString(clientID, 0, 2)
	if IsHZ == "HZ" && len(clientID) == 7 {
		_clientID := SubString(clientID, 2, len(clientID)) //惠州人事存的工号
		theUser, err = GetUserTokenSecretHZ(_clientID, clientSecret)
	} else {
		theUser, err = GetUserTokenSecret(clientID, clientSecret)
	}

	//theUser, err := GetUserTokenSecret(clientID, clientSecret)

	if (err != nil || len(theUser.EmployeeNo) == 0) && (clientID != "00001" && clientID != "99999") {
		err = errors.New("Account does not exist")
		return
	}
	//生成新的令牌前，作废此前的令牌 todo
	gt := oauth2.GrantType("client_credentials")
	tgr := &oauth2.TokenGenerateRequest{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scope:        "admin",
	}

	clientStore := store.NewClientStore()

	clientStore.Set(clientID, &models.Client{
		ID:     clientID,
		Secret: clientSecret,
	})

	TokenManager.MapClientStorage(clientStore)
	ti, err = TokenManager.GenerateAccessToken(gt, tgr)
	// token = ti.GetAccess()

	return
}

type TokenUser struct {
	Id         bson.ObjectId `bson:"_id"`
	EmployeeNo string        `bson:"ID"`
}

func GetUserTokenSecret(clientID string, clientSecret string) (theOne *TokenUser, err error) {
	session, err := mgo.Dial(MONGODB)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("HR").C("EmployeeInfo")
	query := c.Find(bson.M{"ID": clientID, "Info.ePhone": clientSecret, "eWorkState": "1"})
	err = query.One(&theOne)
	return
}

//惠州工业园
func GetUserTokenSecretHZ(clientID string, clientSecret string) (theOne *TokenUser, err error) {
	session, err := mgo.Dial(MONGODB)
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("HZHR").C("EmployeeInfo")
	query := c.Find(bson.M{"ID": clientID, "Info.ePhone": clientSecret, "eWorkState": "1"})
	err = query.One(&theOne)
	return
}

func GetWorkFlowToken(ctx *gin.Context) (token string) {
	AccessToken, _ := ctx.Get("AccessToken")
	p, _ := AccessToken.(*OAuth2Models.Token)
	token = p.Access
	return
}
