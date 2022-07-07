package service

import (
	"context"
	"encoding/json"
	"golang_lopei_grpc_server/repository"
)

// type LopeiService interface{
// 	CheckBalanceMessage(ctx context.Context, in *CheckBalanceMessage)(*ResultMessage, error)

// 	DoPayment(ctx context.Context, in *PaymentMessage)(*ResultMessage, error)
// }

type LopeiService struct {
	repo repository.LopeiRepository
	UnimplementedLopeiPaymentServer
}

func (c *LopeiService) CheckBalanceMessage(ctx context.Context, in *CheckBalanceMessage) (*ResultMessage, error) {
	lopeId := in.LopeiId
	customer, err := c.repo.RetrieveById(lopeId)

	if err != nil {
		return nil, err 
	}

	l, err := json.Marshal(customer) 
	if err != nil {
		return nil, err 
	}
	ResultMessage := &ResultMessage{
		Result: string(l),
		Eror: nil,
	}
	return ResultMessage, nil 
}

func (l *LopeiService) DoPayment(ctx context.Context, in *PaymentMessage)(*ResultMessage, error){
	lopeId := in.LopeId
	amount := in.Amount

	customer, err := l.repo.RetrieveById(lopeId)
	if err != nil {
		return nil, err 
	}

	if customer.Balance < amount {
		return &ResultMessage{
			Result: "FAILED",
			Eror: &Error{
				Code: "X07",
				Message: "Insufficient Balance",
			},
		}, nil 
	}

	resultMessage := &ResultMessage{
		Result: "SUCCESS",
		Eror: nil,
	}
	return resultMessage, nil 
}



func NewLopeiService(repo repository.LopeiRepository) *LopeiService {
	service := new(LopeiService)
	service.repo = repo 
	return service 
}