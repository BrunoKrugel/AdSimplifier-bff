package model

type KiwifyRequest struct {
	OrderID             string      `json:"order_id"`
	OrderRef            string      `json:"order_ref"`
	OrderStatus         string      `json:"order_status"`
	PaymentMethod       string      `json:"payment_method"`
	StoreID             string      `json:"store_id"`
	PaymentMerchantID   string      `json:"payment_merchant_id"`
	Installments        int         `json:"installments"`
	CardType            string      `json:"card_type"`
	CardLast4Digits     string      `json:"card_last4digits"`
	CardRejectionReason interface{} `json:"card_rejection_reason"`
	BoletoURL           interface{} `json:"boleto_URL"`
	BoletoBarcode       interface{} `json:"boleto_barcode"`
	BoletoExpiryDate    interface{} `json:"boleto_expiry_date"`
	PixCode             interface{} `json:"pix_code"`
	PixExpiration       interface{} `json:"pix_expiration"`
	SaleType            string      `json:"sale_type"`
	CreatedAt           string      `json:"created_at"`
	UpdatedAt           string      `json:"updated_at"`
	ApprovedDate        string      `json:"approved_date"`
	RefundedAt          interface{} `json:"refunded_at"`
	Product             struct {
		ProductID   string `json:"product_id"`
		ProductName string `json:"product_name"`
	} `json:"Product"`
	Customer struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Mobile   string `json:"mobile"`
		Cpf      string `json:"CPF"`
		IP       string `json:"ip"`
	} `json:"Customer"`
	Commissions struct {
		ChargeAmount         interface{}  `json:"charge_amount"`
		ProductBasePrice     interface{}  `json:"product_base_price"`
		KiwifyFee            interface{}  `json:"kiwify_fee"`
		CommissionedStores   interface{}  `json:"commissioned_stores"`
		MyCommission         interface{}  `json:"my_commission"`
		FundsStatus          interface{}  `json:"funds_status"`
		EstimatedDepositDate interface{}  `json:"estimated_deposit_date"`
		DepositDate          interface{}  `json:"deposit_date"`
	} `json:"Commissions"`
	TrackingParameters struct {
		Src         string      `json:"src"`
		Sck         string      `json:"sck"`
		UtmSource   interface{} `json:"utm_source"`
		UtmMedium   string      `json:"utm_medium"`
		UtmCampaign string      `json:"utm_campaign"`
		UtmContent  string      `json:"utm_content"`
		UtmTerm     interface{} `json:"utm_term"`
	} `json:"TrackingParameters"`
	AccessURL string `json:"access_url"`
}
