package db

type Shop struct {
	ID   uint   `db:"id"`
	Name string `db:"name"`
}

func ShopsAll() (*[]Shop, error) {
	var shops = make([]Shop, 0)
	err := db.Select(&shops, "SELECT * FROM shops")

	return &shops, err
}
