package MockExample

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"
)

type DatabaseMock struct {
	mock.Mock
}

func (db *DatabaseMock) GetUserName(id int) string {
	args := db.Called(id)
	return args.String(0)
}

func TestPrintUserName(t *testing.T) {
	dbMock := new(DatabaseMock)
	dbMock.On("GetUserName", 1).Return("Alice")

	name := dbMock.GetUserName(1)

	dbMock.AssertExpectations(t)

	fmt.Print(name)
}
