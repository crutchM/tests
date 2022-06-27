package l0

type Order struct {
	Uid               string      `json:"order_uid"`
	Track             string      `json:"track_number"`
	Delivery          Delivery    `json:"delivery"`
	Payment           Payment     `json:"payment"`
	Items             []OrderItem `json:"items"`
	Locale            string      `json:"locale"`
	InternalSignature string      `json:"internal_signature"`
	Customer          string      `json:"customer_id"`
	DeliveryService   string      `json:"delivery_service"`
	ShardKey          string      `json:"shard_key"`
	SmId              int         `json:"sm_id"`
	DateCreated       string      `json:"date_created"`
	OofShard          string      `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
	OrderId string `json:"orderId"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
	OrderId      string `json:"orderId"`
}

type OrderItem struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        int    `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`

	OrderId string `json:"orderId"`
}
