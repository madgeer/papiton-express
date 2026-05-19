package repository

import (
	"database/sql"
	"errors"
)

/*
* PostgresSortingRepo adalah implementasi SortingRepository menggunakan PostgreSQL.
* (Masih kosong/belum diimplementasikan agar testnya failed sesuai requirements)
 */
type PostgresSortingRepo struct {
	db *sql.DB
}

/* NewPostgresSortingRepo membuat instance baru untuk PostgresSortingRepo */
func NewPostgresSortingRepo(db *sql.DB) *PostgresSortingRepo {
	return &PostgresSortingRepo{db: db}
}

func (r *PostgresSortingRepo) AssignToLane(resi string, laneID string) error {
	// TODO: implementasi query INSERT/UPDATE untuk memasukkan paket ke jalur
	return errors.New("database belum diimplementasikan")
}

func (r *PostgresSortingRepo) GetPackagesInLane(laneID string) ([]string, error) {
	// TODO: implementasi query SELECT resi berdasarkan laneID
	return nil, errors.New("database belum diimplementasikan")
}
