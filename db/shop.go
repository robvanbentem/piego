package db

type Shop struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func ShopsAll() ([]Shop, error) {
	var shops []Shop
	err := db.Select(&shops, "select * from shops")

	return shops, err

}
