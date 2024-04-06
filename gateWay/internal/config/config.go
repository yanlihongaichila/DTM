package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	GoodsClient zrpc.RpcClientConf
	OrderClient zrpc.RpcClientConf
	Alpay       Alpay `json:"alpay"`
}

type Alpay struct {
	AppID        string `json:"appID"`
	PrivateKey   string `json:"privateKey"`
	IsProduction bool   `json:"isProduction"`
	PublicKey    string `json:"publicKey"`
	NotifyURL    string `json:"notifyURL"`
	ReturnURL    string `json:"returnURL"`
}
