package request

type CertificateCreateRequest struct {
	UniqueID               string            `json:"unique_id,omitempty"`
	ProductID              string            `json:"product_id,omitempty"`
	Period                 string            `json:"period,omitempty"`
	DomainDCV              map[string]string `json:"domain_dcv,omitempty"`
	CSR                    string            `json:"csr,omitempty"`
	Renew                  int               `json:"renew,omitempty"`
	Organization           string            `json:"organization,omitempty"`
	OrganizationUnit       string            `json:"organization_unit,omitempty"`
	RegisteredAddressLine1 string            `json:"registered_address_line1,omitempty"`
	SerialNo               string            `json:"serial_no,omitempty"`
	Country                string            `json:"country,omitempty"`
	State                  string            `json:"state,omitempty"`
	City                   string            `json:"city,omitempty"`
	PostalCode             string            `json:"postal_code,omitempty"`
	OrganizationPhone      string            `json:"organization_phone,omitempty"`
	DateOfIncorporation    string            `json:"date_of_incorporation,omitempty"`
	ContactName            string            `json:"contact_name,omitempty"`
	ContactTitle           string            `json:"contact_title,omitempty"`
	ContactPhone           string            `json:"contact_phone,omitempty"`
	ContactEmail           string            `json:"contact_email,omitempty"`
	NotifyURL              string            `json:"notify_url,omitempty"`
}
