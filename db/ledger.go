package db

type LedgerEntry struct {
	ID     int64   `db:"id"`
	UserID int64   `db:"user_id"`
	Name   string  `db:"name"`
	Amount float32 `db:"amount"`
	Date   string  `db:"date"`
}

func LedgerAll() (*[]LedgerEntry, error) {
	entries := make([]LedgerEntry, 0)

	err := db.Select(&entries, "SELECT * FROM ledger")
	return &entries, err
}

func LedgerForDate(date string) (*[]LedgerEntry, error) {
	entries := make([]LedgerEntry, 0)

	err := db.Select(&entries, "SELECT * FROM ledger WHERE `date` = ?", date)
	return &entries, err
}
