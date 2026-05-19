package service

/* StatusService khusus menangani pembaruan status paket secara terpusat */
type StatusService struct{}

/* NewStatusService adalah constructor untuk inisialisasi service */
func NewStatusService() *StatusService {
	return &StatusService{} // Dibuat agar bisa compile, akan dikerjakan pada tahap berikutnya.
}

/* UpdatePackageStatus mengupdate status tunggal ke DB */
func (s *StatusService) UpdatePackageStatus(resi string, status string) error {
	// FIXME: Dikosongkan agar test FAILED, akan dikerjakan pada tahap berikutnya.
	return nil
}
