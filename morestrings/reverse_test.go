package morestrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

var serviceMock = &service

func TestReverseRunes(t *testing.T) {

	// override
	var sameValueServiceMock = &SameValueServiceMock{}
	serviceMock.sameValueService = sameValueServiceMock

	//call
	result := ReverseRunes("tteesstt")

	//assert called times
	if sameValueServiceMock.numberOfCalledTimes != 1 {
		t.Errorf("SameValue.Call method should called %v, but %v instead.", 1, 0)
	}
	if result != "ttsseett" {
		t.Errorf("ReverseRunes method should returned %v, but %v instead.", "ttsseett", result)
	}
	fmt.Println(&serviceMock.sameValueService)
}

func TestTestify(t *testing.T) {
	sameValueServiceMockTestify := new(SameValueServiceMockTestify)
	sameValueServiceMockTestify.On("Call", mock.AnythingOfType("string")).Return("testify")

	//override
	serviceMock.sameValueService = sameValueServiceMockTestify
	//call
	result := ReverseRunes("tteesstt")

	//assert call as expected
	sameValueServiceMockTestify.AssertExpectations(t)
	//assert called times
	sameValueServiceMockTestify.AssertNumberOfCalls(t, "Call", 1)
	assert.EqualValues(t, "ttsseett", result, "ReverseRunes method should returned %v, but %v instead.", "ttsseett", result)
}
