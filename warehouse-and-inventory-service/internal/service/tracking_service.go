package service

/* TrackingService menangani komunikasi dengan service Tracking eksternal */
type TrackingService struct{}

/* NewTrackingService adalah constructor untuk inisialisasi service */
func NewTrackingService() *TrackingService {
	return &TrackingService{}
}

/* NotifyTrackingService memberi tahu Tracking Service bahwa status berubah */
func (s *TrackingService) NotifyTrackingService(resi string, status string) error {
	// FIXME: Dikosongkan agar test FAILED, akan dikerjakan pada tahap berikutnya.
	return nil
}

/* FormatManifestData merangkum data manifest untuk dikirim ke sistem eksternal */
func (s *TrackingService) FormatManifestData(manifestID string) (string, error) {
	// FIXME: Dikosongkan agar test FAILED, akan dikerjakan pada tahap berikutnya.
	return "", nil
}
