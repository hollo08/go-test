package test2

import (
	"fmt"
	gomock "github.com/golang/mock/gomock"
	"github.com/hollo08/go-test/testmock"
	"testing"
)

func Test_getUser(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockMyFunc := testmock.NewMockMyFunc(mockCtl)
	mockMyFunc.EXPECT().GetInfo().Return("xiaomotong")
	v := testmock.getUser(mockMyFunc)
	if v == "xiaomotong" {
		fmt.Println("get user right!")
	} else {
		t.Error("get error user")
	}
}
