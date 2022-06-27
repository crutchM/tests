package l0

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

type DataBase struct {
	sync.RWMutex
	db *sqlx.DB
}

func NewDataBase(db *sqlx.DB) *DataBase {
	return &DataBase{db: db}
}

func (s *DataBase) Write(value Order) {
	s.Lock()
	s.db.QueryRow("insert into Order(orderUid, trackNumber, entry, locale, internalSignature, customerId, deliveryService, shardKey, )values ()")
	defer s.Unlock()
}

func (s *DataBase) insertIntoDelivery(delivery Delivery, orderId string) int {
	var id int
	row := s.db.QueryRow("INSERT INTO delivery(name, phone, zip, city, address, region, email, orderID) "+
		"values ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id", delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email, orderId)
	if err := row.Scan(&id); err != nil {
		return 0
	}
	return id
}

func (s *DataBase) insertIntoPayment(payment Payment, orderId string) int {
	var id int
	row := s.db.QueryRow("INSERT INTO payment(transaction, requestId, currency, provider, amount, paymentDt, bank, deliveryCost, goodsTotal, customFee, orderId)"+
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		payment.Transaction, payment.RequestId, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee, orderId)

	if err := row.Scan(&id); err != nil {
		return 0
	}
	return id
}
func (s *DataBase) insertIntoItem(item OrderItem, orderId string) int {
	var id int
	row := s.db.QueryRow("INSERT INTO item(chrtId, trackNumber, price, rid, name, sale, size, totalPrice, nmId, brand, status, orderId)"+
		"values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) RETURNING id",
		item.ChrtId, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status, orderId)
	if err := row.Scan(&id); err != nil {
		return 0
	}
	return id
}
