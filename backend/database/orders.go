package database

func MakeOrder(orders []Order) {
	for _, ord := range orders {
		GetDB().Create(&Order{Name: ord.Name, Quantity: ord.Quantity})
	}
}
