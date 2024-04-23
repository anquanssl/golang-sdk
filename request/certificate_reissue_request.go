package request

// 重签证书请求
type CertificateReissueRequest struct {
    ServiceID              string            `json:"service_id,omitempty"`               // 必传,下单时返回的id
    DomainDCV              map[string]string `json:"domain_dcv,omitempty"`               // 必传
    CSR                    string            `json:"csr,omitempty"`                      // 必传,客户上传的CSR
    Renew                  int               `json:"renew,omitempty"`                    // OV/EV必传,是否为续费订单
    Organization           string            `json:"organization,omitempty"`             // OV/EV必传,公司名称
    OrganizationUnit       string            `json:"organization_unit,omitempty"`        // OV/EV必传,公司部门
    RegisteredAddressLine1 string            `json:"registered_address_line1,omitempty"` // OV/EV必传,公司注册地址
    SerialNo               string            `json:"serial_no,omitempty"`                // OV/EV必传,公司注册号，三证合一
    Country                string            `json:"country,omitempty"`                  // OV/EV必传,2位国别码，大写
    State                  string            `json:"state,omitempty"`                    // OV/EV必传,省份
    City                   string            `json:"city,omitempty"`                     // OV/EV必传,城市
    PostalCode             string            `json:"postal_code,omitempty"`              // OV/EV必传,邮编
    OrganizationPhone      string            `json:"organization_phone,omitempty"`       // OV/EV必传,组织注册登记电话
    DateOfIncorporation    string            `json:"date_of_incorporation,omitempty"`    // OV/EV必传,成立日期
    ContactName            string            `json:"contact_name,omitempty"`             // OV/EV必传,联系人
    ContactTitle           string            `json:"contact_title,omitempty"`            // OV/EV必传,联系人职位
    ContactPhone           string            `json:"contact_phone,omitempty"`            // OV/EV必传,联系人电话
    ContactEmail           string            `json:"contact_email,omitempty"`            // 必传,联系人邮箱
    NotifyURL              string            `json:"notify_url,omitempty"`               // 必传,证书颁发后的通知地址
}
