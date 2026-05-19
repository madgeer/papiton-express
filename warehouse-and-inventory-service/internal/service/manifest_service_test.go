package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
* MockManifestRepo adalah simulasi/database palsu untuk testing
*/
type MockManifestRepo struct {
	mock.Mock
}

/* CreateManifest membuat manifest baru */
func (m *MockManifestRepo) CreateManifest(truckID string, driverName string) (string, error) {
	args := m.Called(truckID, driverName)
	return args.String(0), args.Error(1)
}

/* AddPackageToManifest menambahkan paket ke manifest */
func (m *MockManifestRepo) AddPackageToManifest(manifestID string, resi string) error {
	args := m.Called(manifestID, resi)
	return args.Error(0)
}

/* UpdateManifestStatus mengupdate status manifest */
func (m *MockManifestRepo) UpdateManifestStatus(manifestID string, status string) error {
	args := m.Called(manifestID, status)
	return args.Error(0)
}

/* GetManifestStatus mengembalikan status manifest */
func (m *MockManifestRepo) GetManifestStatus(manifestID string) (string, error) {
	args := m.Called(manifestID)
	return args.String(0), args.Error(1)
}

/* Kumpulan Unit Test untuk Manifest Service
* Semua tes di bawah ini diset EXPECTATION-nya memanggil DB (Mock).
* Karena service aslinya dikosongkan, tes ini AKAN FAILED.
*/

/* TestCreateNewManifest_Unit untuk testing pembuatan manifest baru */
func TestCreateNewManifest_Unit(t *testing.T) {
	mockRepo := new(MockManifestRepo)
	svc := NewManifestService(mockRepo)

	// Ekspektasi: Service HARUS memanggil fungsi CreateManifest di Repository
	mockRepo.On("CreateManifest", "TRK-001", "Budi").Return("MNF-123", nil)

	manifestID, err := svc.CreateNewManifest("TRK-001", "Budi")

	assert.NoError(t, err)
	assert.Equal(t, "MNF-123", manifestID)
	mockRepo.AssertExpectations(t) // Pasti gagal di sini
}

/* TestAddToManifest_Unit untuk testing penambahan paket ke manifest */
func TestAddToManifest_Unit(t *testing.T) {
	mockRepo := new(MockManifestRepo)
	svc := NewManifestService(mockRepo)

	mockRepo.On("AddPackageToManifest", "MNF-123", "RESI-001").Return(nil)

	err := svc.AddToManifest("MNF-123", "RESI-001")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t) // Pasti gagal di sini
}

/* TestFinalizeManifest_Unit untuk testing finalisasi manifest */
func TestFinalizeManifest_Unit(t *testing.T) {
	mockRepo := new(MockManifestRepo)
	svc := NewManifestService(mockRepo)

	mockRepo.On("UpdateManifestStatus", "MNF-123", "FINALIZED").Return(nil)

	err := svc.FinalizeManifest("MNF-123")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t) // Pasti gagal di sini
}

/* TestDepartManifest_Unit untuk testing departure manifest */
func TestDepartManifest_Unit(t *testing.T) {
	mockRepo := new(MockManifestRepo)
	svc := NewManifestService(mockRepo)

	mockRepo.On("UpdateManifestStatus", "MNF-123", "DEPARTED").Return(nil)

	err := svc.DepartManifest("MNF-123")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t) // Pasti gagal di sini
}

/* TestReceiveManifest_Unit untuk testing penerimaan manifest */
func TestReceiveManifest_Unit(t *testing.T) {
	mockRepo := new(MockManifestRepo)
	svc := NewManifestService(mockRepo)

	mockRepo.On("UpdateManifestStatus", "MNF-123", "ARRIVED").Return(nil)

	err := svc.ReceiveManifest("MNF-123", "WH-DESTINATION")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t) // Pasti gagal di sini
}
