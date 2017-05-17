package db

type Item struct {
	ID     uint    `db:"id"`
	ShopID uint    `db:"shop_id"`
	Name   string  `db:"name"`
	Price  float32 `db:"price"`
}

func ItemsAll() (*[]Item, error) {
	items := make([]Item, 0)

	err := db.Select(&items, "SELECT * FROM items")
	if err != nil {
		return &items, err
	}

	return &items, nil
}

func ItemsSearch(shopId int64, s string) (*[]Item, error) {
	items := make([]Item, 0)

	err := db.Select(&items, "SELECT * FROM items WHERE shop_id = ? AND name LIKE ?", shopId, "%"+s+"%")
	if err != nil {
		return &items, err
	}

	return &items, nil
}
