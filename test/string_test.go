package test

import (
	"example/user/hello/morestrings"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func init() {
	fmt.Println("init string_test")
}

type SameValueServiceMock struct {
	numberOfCalledTimes int
}

func (mock *SameValueServiceMock) Call(s string) string {
	mock.numberOfCalledTimes = mock.numberOfCalledTimes + 1
	fmt.Println("standard mock")
	return "mock " + s
}

type SameValueServiceMockTestify struct {
	mock.Mock
}

func (testify *SameValueServiceMockTestify) Call(s string) string {
	fmt.Println("testify mock")
	ret := testify.Called(s)
	return ret.String(0)
}

// func TestMain(m *testing.M) {
// 	fmt.Println("main http")

// 	os.Exit(m.Run())
// }

// func TestMainString(m *testing.M) {
// 	fmt.Println("main http")

// 	os.Exit(m.Run())
// }

var serviceMock = &morestrings.Services

func TestReverseRunes(t *testing.T) {

	// override
	var sameValueServiceMock = &SameValueServiceMock{}
	serviceMock.SameValueService = sameValueServiceMock

	//call
	result := morestrings.ReverseRunes("tteesstt")

	//assert called times
	if sameValueServiceMock.numberOfCalledTimes != 1 {
		t.Errorf("SameValue.Call method should called %v, but %v instead.", 1, 0)
	}
	if result != "ttsseett" {
		t.Errorf("ReverseRunes method should returned %v, but %v instead.", "ttsseett", result)
	}
	fmt.Println(&serviceMock.SameValueService)
}

func TestTestify(t *testing.T) {
	sameValueServiceMockTestify := new(SameValueServiceMockTestify)
	sameValueServiceMockTestify.On("Call", mock.AnythingOfType("string")).Return("testify")

	//override
	serviceMock.SameValueService = sameValueServiceMockTestify
	//call
	result := morestrings.ReverseRunes("tteesstt")

	//assert call as expected
	sameValueServiceMockTestify.AssertExpectations(t)
	//assert called times
	sameValueServiceMockTestify.AssertNumberOfCalls(t, "Call", 1)
	assert.EqualValues(t, "ttsseett", result, "ReverseRunes method should returned %v, but %v instead.", "ttsseett", result)
}

// func TestHttp(t *testing.T) {
// }
