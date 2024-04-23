package resource

import (
	"encoding/json"

	anquanssl "github.com/anquanssl/golang-sdk"
	"github.com/anquanssl/golang-sdk/request"
)

type Order struct {
    client *anquanssl.Client
}

func NewOrder(client *anquanssl.Client) *Order {
    return &Order{
        client: client,
    }
}

func structToMap(data interface{}) map[string]interface{} {
    var m map[string]interface{}
    ms, _ := json.Marshal(data)
    json.Unmarshal(ms, &m)
    return m
}

func structToStringMap(data interface{}) map[string]string {
    var m map[string]string
    ms, _ := json.Marshal(data)
    json.Unmarshal(ms, &m)
    return m
}

func (o *Order) CertificateCreate(certificateCreateRequest request.CertificateCreateRequest) (map[string]interface{}, error) {
    return o.client.Post("/certificate/create", make(map[string]string), structToMap(certificateCreateRequest))
}

func (o *Order) CertificateDetail(certificateDetailRequest request.CertificateDetailRequest) (map[string]interface{}, error) {
    return o.client.Get("/certificate/detail", structToStringMap(certificateDetailRequest), make(map[string]interface{}))
}

func (o *Order) CertificateReissue(certificateReissueRequest request.CertificateReissueRequest) (map[string]interface{}, error) {
    return o.client.Post("/certificate/reissue", make(map[string]string), structToMap(certificateReissueRequest))
}

func (o *Order) CertificateValidateDCV(certificateValidateDCVRequest request.CertificateValidateDCVRequest) (map[string]interface{}, error) {
    return o.client.Post("/certificate/validate-dcv", make(map[string]string), structToMap(certificateValidateDCVRequest))
}

func (o *Order) CertificateRefund(certificateRefundRequest request.CertificateRefundRequest) (map[string]interface{}, error) {
    return o.client.Post("/certificate/refund", make(map[string]string), structToMap(certificateRefundRequest))
}
