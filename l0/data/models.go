package data

type Order struct {
	Uid               string      `json:"order_uid" db:"orderUid"`
	Track             string      `json:"track_number" db:"trackNumber"`
	Entry             string      `json:"entry" db:"entry"`
	Delivery          Delivery    `json:"delivery" db:"locale"`
	Payment           Payment     `json:"payment" db:""`
	Items             []OrderItem `json:"items" db:""`
	Locale            string      `json:"locale" db:""`
	InternalSignature string      `json:"internal_signature" db:"internalSignature"`
	Customer          string      `json:"customer_id" db:"customerId"`
	DeliveryService   string      `json:"delivery_service" db:"deliveryService"`
	ShardKey          string      `json:"shard_key" db:"shardKey"`
	SmId              int         `json:"sm_id" db:"smId"`
	DateCreated       string      `json:"date_created" db:"dateCreated"`
	OofShard          string      `json:"oof_shard" db:"oofShard"`
}

type Delivery struct {
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Zip     string `json:"zip" db:"zip"`
	City    string `json:"city" db:"city"`
	Address string `json:"address" db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email" db:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction" db:"trans"`
	RequestId    string `json:"request_id" db:"requestId"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int    `json:"amount" db:"amount"`
	PaymentDt    int    `json:"payment_dt" db:"paymentDt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int    `json:"delivery_cost" db:"deliveryCost"`
	GoodsTotal   int    `json:"goods_total" db:"goodsTotal"`
	CustomFee    int    `json:"custom_fee" db:"customFee"`
}

type OrderItem struct {
	ChrtId      int    `json:"chrt_id" db:"chrtId"`
	TrackNumber string `json:"track_number" db:"trackNumber"`
	Price       int    `json:"price" db:"price"`
	Rid         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int    `json:"sale" db:"sale"`
	Size        int    `json:"size" db:"si"`
	TotalPrice  int    `json:"total_price" db:"totalPrice"`
	NmId        int    `json:"nm_id" db:"nmId"`
	Brand       string `json:"brand" db:"brand"`
	Status      int    `json:"status" db:"status"`
}
