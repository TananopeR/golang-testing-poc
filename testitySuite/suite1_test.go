package testitySuite

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type suite1 struct {
	suite.Suite
}

type suite2 struct {
	suite.Suite
}

func (suite *suite1) SetupTest() {
	fmt.Println("suite 1 setup")
}
func (suite *suite2) SetupTest() {
	fmt.Println("suite 2 setup")
}

func (suite *suite1) TestSuite1_1() {
	fmt.Println("suite 1_1 test")
}

func (suite *suite1) TestSuite1_2() {
	fmt.Println("suite 1_2 test")
}

func (suite *suite2) TestSuite2_1() {
	fmt.Println("suite 2_1 test")
}

func (suite *suite2) TestSuite2_2() {
	fmt.Println("suite 2_2 test")
}

func TestExampleTestSuite1(t *testing.T) {
	suite.Run(t, new(suite1))
}

func TestExampleTestSuite2(t *testing.T) {
	suite.Run(t, new(suite2))
}
