package personal_finance

import (
	"errors"
	"time"
)

type Record struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type IPersonalFinanceService interface {
	getAllRecords(f any, l ...int) (*[]any, error)
	getRecordById(id string) (*any, error)
	updateRecord(r any) (*any, error)
	deleteRecords(id []string) ([]bool, error)
}

type PersonalFinanceService struct {
	Name string
}

var _ IPersonalFinanceService = &PersonalFinanceService{}

func NewPersonalFinanceService() {
}

func (pf PersonalFinanceService) getAllRecords(f any, l ...int) (*[]any, error) {
	return nil, errors.New("this function is incomplete")
}

func (pf PersonalFinanceService) getRecordById(id string) (*any, error) {
	return nil, errors.New("this function is incomplete")
}

func (pf PersonalFinanceService) updateRecord(r any) (*any, error) {
	return nil, errors.New("this function is incomplete")
}

func (pf PersonalFinanceService) deleteRecords(ids []string) ([]bool, error) {
	return nil, errors.New("this function is incomplete")
}
