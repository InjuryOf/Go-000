package main

import (
	"database/sql"
	"errors"
	"fmt"
	errors2 "github.com/pkg/errors"
)

// 服务层调用
func main() {
	orderNo := "Go20201201"
	err :=  orderDetails(orderNo)

	// 1、 在最外层调用时进行判定，如果是数据集为空，则返回给服务API调用者一个相应的错误码，用于前端展示处理。
	// 2、 如果是其他异常，则打印错误信息，排查错误
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("data set is null code 3001", )
	} else {
		fmt.Printf("service : %+v\n", err)
	}
}

type Order struct {
	orderNo string // 订单编号
	channel string // 订单来源
}

// 数据层（dao-OrderModel）
func findOne(orderNo string) (*Order, error) {
	fmt.Printf("search params： %s\n", orderNo)
	return nil, errors2.Wrap(sql.ErrNoRows, "SQL查询数据集为空")
	//return nil, errors2.Wrap(sql.ErrConnDone, "链接异常")
}

// 业务逻辑层(biz)
func orderDetails(orderNo string) error {
	_, error := findOne(orderNo)
	return error
}
