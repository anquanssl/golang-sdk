[<p align="center"><img src="https://github.com/anquanssl/.github/raw/main/profile/logo_dark.png" width="600" height="85"/></p>](https://www.anquanssl.com?__utm_from=github-org-profile#gh-dark-mode-only)
[<p align="center"><img src="https://github.com/anquanssl/.github/raw/main/profile/logo_light.png" width="600" height="85"/></p>](https://www.anquanssl.com?__utm_from=github-org-profile#gh-light-mode-only)

## AnquanSSL

AnquanSSL, aka "Security SSL", also known as "安全 SSL" in Mandarin, founded in 2018, and our mission is providing affordable, secure, and enhanced TLS utilization experiences in the Greater China market.

这是 [安全SSL](https://www.anquanssl.com) 开放API的 Golang SDK.

[获取](https://www.anquanssl.com/dashboard/api-credentials) `AccessKey` 秘钥对.

此SDK包仅面向开发者提供支持，若您是分销商，您可以需要:
- [AnquanSSL Module for WHMCS]()
- [AnquanSSL Module for IDCSmart]()

如果您要其它编程语言的开发者，您可能需要
- [AnquanSSL PHP SDK](https://github.com/anquanssl/sdk)
- [AnquanSSL Python SDK](https://github.com/anquanssl/python-sdk)
- [AnquanSSL NodeJS SDK](https://github.com/anquanssl/nodejs-sdk)
- [AnquanSSL Golang SDK](https://github.com/anquanssl/golang-sdk)
- [AnquanSSL Java SDK](https://github.com/anquanssl/java-sdk)


## 安装

```bash
go install https://github.com/anquanssl/golang-sdk
```

## 使用

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	anquanssl "github.com/anquanssl/golang-sdk"
	request "github.com/anquanssl/golang-sdk/request"
	resource "github.com/anquanssl/golang-sdk/resource"
)

func main() {
	accessKeyID := ""
	accessKeySecrte := ""

	client := anquanssl.NewClient(accessKeyID, accessKeySecrte, "")
	// product := resource.NewProduct(client)
	order := resource.NewOrder(client)
	var resp map[string]interface{}

	// resp, _ = product.ProductList()
	// fmt.Println("product list:", resp)

	fmt.Println("=======")

	certificateCreateRequest := request.CertificateCreateRequest{}
	certificateCreateRequest.ProductID = "sslcom_dv_flex"
	certificateCreateRequest.Period = "annually"
	certificateCreateRequest.CSR = `-----BEGIN CERTIFICATE REQUEST-----
MIICsTCCAZkCAQAwQjELMAkGA1UEBhMCQ04xDTALBgNVBAgTBHRlc3QxDTALBgNV
BAcTBHRlc3QxFTATBgNVBAMMDCouZG9tYWluLmNvbTCCASIwDQYJKoZIhvcNAQEB
BQADggEPADCCAQoCggEBAKvgoSs6HahR2RACS2j+hjMtcCUKMWW0kpB96JM3ITEp
OdmkbsuFJfRMavycNvwbvk21D/gq32YJ76Tp1zNEZh79bU/6gwMITYzjMdpguUj5
481cxJIFj/uCnTXdRBOT4ZhsO9oyf42HAnGbo7ZWIbzkQpOeKio9ytAR3JHmeyj5
eVnIfs7rE0poeGR27Kbq74um4dFMj1C8GvkXud6AFqFTJX+xQmWloZeROGfzK+bE
D4rW6olqxkR82cYYQRW1lpk5WHXlYwICiAVhPg4fcQY2RhOxQDIOY5Rio3+oRX/k
QttRNyOPmQAp2J8Pp48IEJg1W5QG7fFo2iCmIVY1SH8CAwEAAaAqMCgGCSqGSIb3
DQEJDjEbMBkwFwYDVR0RBBAwDoIMKi5kb21haW4uY29tMA0GCSqGSIb3DQEBCwUA
A4IBAQAAQq87dAn0YsAkpfvmQSZUov27Qp3M/8JluLv0ykDp4aIuB7uFcU2cgzuK
Zqn7DtrLnjsGmaiiEvsOje+6YH6Skcn/kACGDFg3PT3+Zd/Ar0SZsBZVSzAaI7t0
bc1hsIW0PtbDIAUmwcJEWzCxBlHtANWhnmEexbubAEzbM9BPLYN/s7JluOFDxftP
lXp2j+c5/l8E0huwqnGgSnG/g75zRM4sDjKHmWYMapVaHtFzYktn5fklJxCIOmkY
d+ZnVcZEq5UjeEgUG9P0WguwvwZe0szM8ae+cMxJ/mDcrt/g7ammTD80XfQImYij
aKJmfyDcygIdCZ6uk87LmN0UF4rd
-----END CERTIFICATE REQUEST-----`

	// generate random
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
    UniqueID := string(b)

	certificateCreateRequest.UniqueID = UniqueID
	certificateCreateRequest.ContactEmail = "email@example.org"
	certificateCreateRequest.ContactName = "~"
	certificateCreateRequest.DomainDCV = make(map[string]string)
	certificateCreateRequest.DomainDCV["*.domain.com"] = "dns"
	certificateCreateRequest.DomainDCV["domain.com"] = "dns"
	certificateCreateRequest.DomainDCV["*.domain2.com"] = "dns"
	certificateCreateRequest.DomainDCV["domain2.com"] = "dns"
	certificateCreateRequest.NotifyURL = "https://my-callback.app/notify-url"

	resp, _ = order.CertificateCreate(certificateCreateRequest)

	// get `service_id` from map `resp`
	serviceID := resp["data"].(map[string]interface{})["service_id"].(string)
	fmt.Printf("certificateCreate serviceID := ", serviceID)
	fmt.Printf("\n")

	certificateDetailRequest := request.CertificateDetailRequest{}
	certificateDetailRequest.ServiceID = serviceID
	resp, _ = order.CertificateDetail(certificateDetailRequest)
	fmt.Println("certificateDetailRequest:", resp)
	fmt.Printf("\n")

	certificateValidateDCVRequest := request.CertificateValidateDCVRequest{}
	certificateValidateDCVRequest.ServiceID = serviceID
	resp, _ = order.CertificateValidateDCV(certificateValidateDCVRequest)
	fmt.Println("certificateValidDCV:", resp)
	fmt.Printf("\n")

	certificateRefundRequest := request.CertificateRefundRequest{}
	certificateRefundRequest.ServiceID = serviceID
	resp, _ = order.CertificateRefund(certificateRefundRequest)
	fmt.Println("certificateRefund:", resp)
	fmt.Printf("\n")
}
```

## 贡献

特别鸣谢以下工程师对本项目的贡献:

[@jellnicy](https://github.com/jellnicy)