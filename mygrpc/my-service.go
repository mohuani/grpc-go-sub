package mygrpc

import (
	"context"
	"fmt"
)

type MyService struct{}

func (s *MyService) Sub(ctx context.Context, req *SubReq) (*SubResp, error) {
	fmt.Println(req)
	number := subTwoNumbers(req.Number1, req.Number2)
	return &SubResp{
		Number: number,
	}, nil
}

func subTwoNumbers(number1, number2 int64) int64 {
	return number1 - number2
}
