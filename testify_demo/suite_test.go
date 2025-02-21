package testify_demo

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestExampleTestSuite(t *testing.T) {
	suiteToTest := &ExampleTestSuite{}
	suite.Run(t, suiteToTest)
}

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type ExampleTestSuite struct {
	suite.Suite
	Value int
}

// SetupTest runs before each test
func (s *ExampleTestSuite) SetupTest() {
	s.Value = 5
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (s *ExampleTestSuite) TestExample1() {
	// assert.Equal(s.T(), 5, s.Value)
	s.Equal(5, s.Value)
}

func (s *ExampleTestSuite) TestExample2() {
	// assert.NotNil(s.T(), s.Value)
	s.NotNil(s.Value)
}
