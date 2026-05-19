package repository

import (
	"database/sql"
	"errors"
)

/*
* PostgresInboundRepo adalah implementasi InboundRepository menggunakan PostgreSQL.
* (Masih kosong/belum diimplementasikan agar testnya failed sesuai requirements)
 */
type PostgresInboundRepo struct {
	db *sql.DB
}

/* NewPostgresInboundRepo membuat instance baru untuk PostgresInboundRepo */
func NewPostgresInboundRepo(db *sql.DB) *PostgresInboundRepo {
	return &PostgresInboundRepo{db: db}
}

func (r *PostgresInboundRepo) UpdateStockStatus(resi string, warehouseID string, status string) error {
	// TODO: implementasi query update
	return errors.New("database belum diimplementasikan")
}

func (r *PostgresInboundRepo) GetItemByResi(resi string) (string, error) {
	// TODO: implementasi query select
	return "", errors.New("database belum diimplementasikan")
}

func (r *PostgresInboundRepo) UpdatePackageMetadata(resi string, instructions []string) error {
	// TODO: implementasi query update
	return errors.New("database belum diimplementasikan")
}
