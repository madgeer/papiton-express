package service

/*
* import package yang dibutuhkan
* testing : package untuk melakukan unit testing
* assert : package untuk melakukan assertion
* mock : package untuk melakukan mocking
*/
import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
* buat simulasi DB
*/
type MockRepo struct {
	mock.Mock
}

/*
* method UpdateStockStatus untuk simulasi DB
*/
func (m *MockRepo) UpdateStockStatus(resi, wh, status string) error {
	args := m.Called(resi, wh, status)
	return args.Error(0)
}

/*
* method GetItemByResi untuk simulasi DB
*/
func (m *MockRepo) GetItemByResi(resi string) (string, error) {
	args := m.Called(resi)
	return args.String(0), args.Error(1)
}

/*
* method UpdatePackageMetadata untuk simulasi DB
*/
func (m *MockRepo) UpdatePackageMetadata(resi string, instructions []string) error {
	args := m.Called(resi, instructions)
	return args.Error(0)
}

/*
* unit test untuk InboundProcess
*/
func TestProcessInbound_Unit(t *testing.T) {
	mockRepo := new(MockRepo) // memanggil struct mockRepo yang sudah dibuat
	svc := NewInboundService(mockRepo) // memanggil struct InboundService dengan mockRepo sebagai argumen

	// Contoh resi dengan format: [kode kota]YYMMDDHHMMSS + 4 char random
	// BDO = Bandung, 240430120000 = YYMMDDHHMMSS, X1Y2 = 4 char random
	resi := "BDO240430120000X1Y2"

	// Ekspektasi: Fungsi UpdateStockStatus HARUS dipanggil oleh service
	mockRepo.On("UpdateStockStatus", resi, "WH-UPI", "AT_HUB").Return(nil)

	// Jalankan fungsi
	err := svc.ProcessInbound(resi, "WH-UPI")

	// ASSERTION: Pasti FAIL karena ProcessInbound di inbound.go belum manggil repo
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

/*
* unit test untuk ValidatePackage
*/
func TestValidatePackage_Unit(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := NewInboundService(mockRepo)

	isValid, isExpress, instructions := svc.ValidatePackage("BDO240430120000X1Y2")
	
	// Secara default return true, false, []
	assert.True(t, isValid)
	assert.False(t, isExpress)
	assert.Empty(t, instructions)
}

/*
* unit test untuk AssignStorageZone. AssignStorageZone itu function untuk menentukan area penyimpanan sementara berdasarkan prioritas paket.
* contohnya paket express di prioritize dan disimpan di ZONE_EXPRESS, sedangkan paket reguler di prioritize dan disimpan di ZONE_REGULAR
*/
func TestAssignStorageZone_Unit(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := NewInboundService(mockRepo)

	// Test untuk paket Express
	zoneExpress := svc.AssignStorageZone("BDO240430120000X1Y2", true)
	assert.Equal(t, "ZONE_EXPRESS", zoneExpress)

	// Test untuk paket Reguler
	zoneReguler := svc.AssignStorageZone("BDO240430120000X1Y2", false)
	assert.Equal(t, "ZONE_REGULAR", zoneReguler)
}

/*
* unit test untuk ApplySpecialHandling
*/
func TestApplySpecialHandling_Unit(t *testing.T) {
	mockRepo := new(MockRepo)
	svc := NewInboundService(mockRepo)

	// Skenario 1: Tanpa instruksi khusus
	err := svc.ApplySpecialHandling("BDO240430120000X1Y2", []string{})
	assert.NoError(t, err)

	// Skenario 2: Dengan instruksi khusus (misal "FRAGILE")
	// Tambahkan ekspektasi agar test failed (karena implementasi belum memanggil repo)
	mockRepo.On("UpdatePackageMetadata", "BDO240430120000X1Y2", []string{"FRAGILE"}).Return(nil)
	err = svc.ApplySpecialHandling("BDO240430120000X1Y2", []string{"FRAGILE"})
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}