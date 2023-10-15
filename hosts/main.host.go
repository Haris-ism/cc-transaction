package host

import (
	"cc-transaction/hosts/callback"
	"cc-transaction/hosts/merchant"
)

type (
	host struct{
		merchant	merchant.MerchantInterface
		callback	callback.CallbackInterface
	}
	HostInterface interface{
		Merchant()merchant.MerchantInterface
		Callback()callback.CallbackInterface
	}
)

func InitHost(merchant merchant.MerchantInterface,callback callback.CallbackInterface) HostInterface {
	return &host{
		merchant: merchant,
		callback:callback,
	}
}

func (h *host) Merchant()merchant.MerchantInterface{
	return h.merchant
}
func (h *host) Callback()callback.CallbackInterface{
	return h.callback
}