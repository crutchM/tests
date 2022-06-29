package db

import (
	"github.com/jmoiron/sqlx"
	"sync"
	"tests/l0/data"
)

type DataBase struct {
	sync.RWMutex
	db *sqlx.DB
}

func NewDataBase(db *sqlx.DB) *DataBase {
	return &DataBase{db: db}
}

func (s *DataBase) Write(value data.Order) {
	s.Lock()
	s.db.QueryRow("insert into order(orderUid, trackNumber,  entry, locale, internalSignature, customerId, deliveryService, shardKey, smId, dateCreated, oofShard)"+
		"values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", value.Uid, value.Track, value.Entry, value.Locale,
		value.InternalSignature, value.Customer, value.DeliveryService, value.ShardKey, value.SmId, value.DateCreated, value.OofShard)
	s.insertIntoPayment(value.Payment, value.Uid)
	s.insertIntoDelivery(value.Delivery, value.Uid)
	for _, val := range value.Items {
		s.insertIntoItem(val, value.Uid)
	}
	defer s.Unlock()
}

func (s *DataBase) insertIntoDelivery(delivery data.Delivery, orderId string) int {
	var id int
	row := s.db.QueryRow("INSERT INTO delivery(name, phone, zip, city, address, region, email, orderId) "+
		"values ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id", delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email, orderId)
	if err := row.Scan(&id); err != nil {
		return 0
	}
	return id
}

func (s *DataBase) insertIntoPayment(payment data.Payment, orderId string) int {
	var id int
	row := s.db.QueryRow("INSERT INTO payment(trans, requestId, currency, provider, amount, paymentDt, bank, deliveryCost, goodsTotal, customFee, orderId)"+
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		payment.Transaction, payment.RequestId, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee, orderId)

	if err := row.Scan(&id); err != nil {
		return 0
	}
	return id
}
func (s *DataBase) insertIntoItem(item data.OrderItem, orderId string) int {
	var id int
	row := s.db.QueryRow("INSERT INTO item(chrtId, trackNumber, price, rid, name, sale, si, totalPrice, nmId, brand, status, orderId)"+
		"values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id",
		item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status, orderId)
	if err := row.Scan(&id); err != nil {
		return 0
	}
	return id
}

func (s *DataBase) GetRow(id string) data.Order {
	s.Lock()
	defer s.Unlock()
	var order data.Order
	var delivery data.Delivery
	var payment data.Payment
	var items []data.OrderItem
	err := s.db.Get(&order, "select orderUid, trackNumber,  entry, locale, internalSignature, customerId, deliveryService, shardKey, smId, dateCreated, oofShard from order where orderUid=$1", id)
	if err != nil {
		return data.Order{}
	}
	err = s.db.Get(&delivery, "select name, phone, zip, city, address, region, email from delivery where orderId=$1", id)
	if err != nil {
		return data.Order{}
	}
	err = s.db.Get(&payment, "select trans, requestId, currency, provider, amount, paymentDt, bank, deliveryCost, goodsTotal, customFee from payment where orderId=$1", id)

	if err != nil {
		return data.Order{}
	}
	err = s.db.Select(&items, "select chrtId, trackNumber, price, rid, name, sale, si, totalPrice, nmId, brand, status from item where orderId=$1", id)
	if err != nil {
		return data.Order{}
	}

	order.Delivery = delivery
	order.Payment = payment
	order.Items = items
	return order
}

func (s *DataBase) GetAll() []data.Order {
	s.Lock()
	defer s.Unlock()
	var orders []data.Order
	var result []data.Order
	err := s.db.Select(&orders, "select orderUid, trackNumber,  entry, locale, internalSignature, customerId, deliveryService, shardKey, smId, dateCreated, oofShard from order where uid=$1")
	if err != nil {
		return nil
	}
	for _, val := range orders {
		result = append(result, s.GetRow(val.Uid))
	}
	return result
}
