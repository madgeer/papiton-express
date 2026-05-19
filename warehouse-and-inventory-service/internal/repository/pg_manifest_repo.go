package repository

import (
	"database/sql"
	"errors"
)

/*
* PostgresManifestRepo adalah implementasi ManifestRepository menggunakan PostgreSQL.
* (Masih kosong/belum diimplementasikan agar testnya failed sesuai requirements)
 */
type PostgresManifestRepo struct {
	db *sql.DB
}

/* NewPostgresManifestRepo membuat instance baru untuk PostgresManifestRepo */
func NewPostgresManifestRepo(db *sql.DB) *PostgresManifestRepo {
	return &PostgresManifestRepo{db: db}
}

func (r *PostgresManifestRepo) CreateManifest(truckID string, driverName string) (string, error) {
	// TODO: implementasi query INSERT INTO manifests
	return "", errors.New("database belum diimplementasikan")
}

func (r *PostgresManifestRepo) AddPackageToManifest(manifestID string, resi string) error {
	// TODO: implementasi query INSERT INTO manifest_items
	return errors.New("database belum diimplementasikan")
}

func (r *PostgresManifestRepo) UpdateManifestStatus(manifestID string, status string) error {
	// TODO: implementasi query UPDATE manifests SET status = ...
	return errors.New("database belum diimplementasikan")
}

func (r *PostgresManifestRepo) GetManifestStatus(manifestID string) (string, error) {
	// TODO: implementasi query SELECT status FROM manifests
	return "", errors.New("database belum diimplementasikan")
}
