package pkg

import (
	"gateWay/internal/svc"
	"github.com/smartwalle/alipay/v3"
	"net/url"
)

// 创建连接
func getClient(svcCtx *svc.ServiceContext) (*alipay.Client, error) {
	alpayConf := svcCtx.Config.Alpay
	var client, err = alipay.New(alpayConf.AppID, alpayConf.PrivateKey, alpayConf.IsProduction)
	if err != nil {
		return nil, err
	}
	err = client.LoadAliPayPublicKey(alpayConf.PublicKey)
	if err != nil {
		return nil, err
	}
	return client, err

}

func GetWebPayUrl(svcCtx *svc.ServiceContext, title, orderNO, amount string) (string, error) {
	alpayConf := svcCtx.Config.Alpay
	var p = alipay.TradePagePay{
		Trade: alipay.Trade{
			NotifyURL:   alpayConf.NotifyURL,
			ReturnURL:   alpayConf.ReturnURL,
			Subject:     title,
			OutTradeNo:  orderNO,
			TotalAmount: amount,
			ProductCode: "FAST_INSTANT_TRADE_PAY",
		},
	}
	cli, err := getClient(svcCtx)
	if err != nil {
		return "", err
	}

	payUrl, err := cli.TradePagePay(p)
	if err != nil {
		return "", err
	}
	return payUrl.String(), nil
}

func VerifySign(svcCtx *svc.ServiceContext, values url.Values) (*alipay.Notification, error) {
	cli, err := getClient(svcCtx)
	if err != nil {
		return nil, err
	}
	err = cli.VerifySign(values)
	return cli.DecodeNotification(values)
}
