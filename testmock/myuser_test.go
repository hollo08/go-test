package testmock

import (
	"fmt"
	gomock "github.com/golang/mock/gomock"
	"github.com/hollo08/go-test/testmock/mock_testmock"
	"testing"
)

func Test_getUser(t *testing.T) {
	mockCtl := gomock.NewController(t)
	mockMyFunc := NewMockMyFunc(mockCtl)
	mockMyFunc.EXPECT().GetInfo().Return("xiaomotong")
	v := getUser(mockMyFunc)
	if v == "xiaomotong" {
		fmt.Println("get user right!")
	} else {
		t.Error("get error user")
	}
}
