package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

/*
* MockSortingRepo adalah simulasi/database palsu untuk testing Sorting
*/
type MockSortingRepo struct {
	mock.Mock
}

/* MockSortingRepo.AssignToLane untuk testing assign lane */
func (m *MockSortingRepo) AssignToLane(resi string, laneID string) error {
	args := m.Called(resi, laneID)
	return args.Error(0)
}

/* MockSortingRepo.GetPackagesInLane untuk testing get packages in lane */
func (m *MockSortingRepo) GetPackagesInLane(laneID string) ([]string, error) {
	args := m.Called(laneID)
	// Type assertion yang aman untuk array/slice
	if val := args.Get(0); val != nil {
		return val.([]string), args.Error(1)
	}
	return nil, args.Error(1)
}

/*
* Kumpulan Unit Test untuk Sorting Service.
* Ekspektasi dibuat, tapi fungsi asli dikosongkan agar FAILED.
*/

/* TestAssignPackageToLane_Unit untuk testing assign lane */
func TestAssignPackageToLane_Unit(t *testing.T) {
	mockRepo := new(MockSortingRepo)
	svc := NewSortingService(mockRepo)

	// Ekspektasi: Service HARUS menyimpan data ke DB (Mock)
	mockRepo.On("AssignToLane", "BDO240430120000X1Y2", "LANE-JAKARTA").Return(nil)

	err := svc.AssignPackageToLane("BDO240430120000X1Y2", "LANE-JAKARTA")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t) // Pasti FAILED
}

/* TestGetPackagesInLane_Unit untuk testing get packages in lane */
func TestGetPackagesInLane_Unit(t *testing.T) {
	mockRepo := new(MockSortingRepo)
	svc := NewSortingService(mockRepo)

	// Persiapan data palsu
	dummyResiList := []string{"RESI-01", "RESI-02"}

	// Ekspektasi: Service mengambil data dari DB
	mockRepo.On("GetPackagesInLane", "LANE-JAKARTA").Return(dummyResiList, nil)

	pkgs, err := svc.GetPackagesInLane("LANE-JAKARTA")

	assert.NoError(t, err)
	assert.Len(t, pkgs, 2)
	mockRepo.AssertExpectations(t) // Pasti FAILED
}
