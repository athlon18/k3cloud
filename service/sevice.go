package service

import (
	"fmt"
	_struct "github.com/athlon18/k3cloud/service/struct"
	"github.com/athlon18/k3cloud/service/struct/response"
	"github.com/athlon18/k3cloud/service/util"
)

type K3config struct {
	CloudUrl   string `json:"cloudUrl"` // 地址
	AcctID     string `json:"acctID"`   //Id
	Username   string `json:"username"`
	Password   string `json:"password"`
	LcID       int    `json:"lcid"`
	session    *util.Browser
	k3Response response.K3ResponseStruct
}

// 初始化
func K3configInit(url string, acctID string, username string, password string) *K3config {
	k3config := &K3config{}
	k3config.CloudUrl = url
	k3config.AcctID = acctID
	k3config.Password = password
	k3config.Username = username
	k3config.LcID = 2052
	k3config.session = util.NewBrowser()
	return k3config
}

/**
K3cloud 登录
*/
func (k3config *K3config) Login() *K3config {
	formParams := util.CreateLoginPostData(k3config.AcctID, k3config.Username, k3config.Password, k3config.LcID)
	res := k3config.session.PostJson(k3config.CloudUrl+util.LOGIN_API, formParams)
	k3Response := response.K3LoginResponseToStruct(res)
	if k3Response.LoginResultType == 0 {
		panic(k3Response.Message)
	}
	return k3config
}

/**
凭证写入
*/
func (k3config *K3config) Voucher(root _struct.Root) *K3config {

	//root.Model
	formParams := util.CreateBusinessPostData("GL_VOUCHER", util.Struct2Map(root))
	res := k3config.session.PostJson(k3config.CloudUrl+util.SAVE_API, formParams)
	k3config.k3Response = response.K3ResponseToStruct(res)
	if k3config.k3Response.Result.ResponseStatus.IsSuccess == false {
		panic(k3config.k3Response.Result.ResponseStatus.Errors)
	}
	return k3config
}

/**
k3cloud 打印
*/
func (k3config *K3config) Print() {
	fmt.Println(k3config.k3Response)
}

/**
返回接口数据
*/
func (k3config *K3config) Get() response.K3ResponseStruct {
	return k3config.k3Response
}

/**
返回API接口数据
*/
func (k3config *K3config) GetResponse() response.K3ResponseStruct {
	return k3config.k3Response
}
