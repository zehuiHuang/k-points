package business

type Info struct {
	UserId       string `json:"user_id"`
	ProductName  string `json:"product_name"`
	BillingType  int    `json:"billing_type"`   //权益类型
	CostTypeCode int    `json:"cost_type_code"` //成本分类码
	RespChan     chan RespPayload
	ExecuteCount int `json:"execute_count"`
}
