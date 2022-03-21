package testitySuite

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type suite3 struct {
	suite.Suite
}

type suite4 struct {
	suite.Suite
}

func (suite *suite3) SetupTest() {
	fmt.Println("suite 3 setup")
}
func (suite *suite4) SetupTest() {
	fmt.Println("suite 4 setup")
}

func (suite *suite3) TestSuite3_1() {
	fmt.Println("suite 3_1 test")
}

func (suite *suite3) TestSuite3_2() {
	fmt.Println("suite 3_2 test")
}

func (suite *suite4) TestSuite4_1() {
	fmt.Println("suite 4_1 test")
}

func (suite *suite4) TestSuite4_2() {
	fmt.Println("suite 4_2 test")
}

func TestExampleTestSuite3(t *testing.T) {
	suite.Run(t, new(suite3))
}

func TestExampleTestSuite4(t *testing.T) {
	suite.Run(t, new(suite4))
}
