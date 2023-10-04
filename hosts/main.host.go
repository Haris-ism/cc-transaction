package host

import "cc-transaction/hosts/merchant"

type (
	host struct{
		merchant	merchant.MerchantInterface
	}
	HostInterface interface{
		Merchant()merchant.MerchantInterface
	}
)

func InitHost(merchant merchant.MerchantInterface) HostInterface {
	return &host{
		merchant: merchant,
	}
}

func (h *host) Merchant()merchant.MerchantInterface{
	return h.merchant
}