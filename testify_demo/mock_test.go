package testify_demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

type ActorI interface {
	DoSomething(number int) (bool, error)
}

type MyObject struct {
	actor ActorI
}

// === Mock object begin ===

// In a real project, the mock object should be auto-generated and put in a separate file
type MockedActor struct {
	mock.Mock
}

// DoSomething is a method on MockedActor that implements some interface
// and just records the activity, and returns what the Mock object tells it to.
//
// In the real object, this method would do something useful, but since this
// is a mocked object - we're just going to stub it out.
//
// NOTE: This method is not being tested here, code that uses this object is.
func (m *MockedActor) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

// === Mock object end ===

// Actual test functions
func TestMyObj_whenActorFails_shouldFails(t *testing.T) {
	// Arrange
	mockActor := &MockedActor{}
	testObj := &MyObject{actor: mockActor}

	// Mock
	mockActor.On("DoSomething", mock.Anything).Return(false, nil)

	// Act
	res := targetFuncThatDoesSomethingWithObj(testObj, 123)

	// Assert
	mockActor.AssertExpectations(t)
	assert.False(t, res)
}

func targetFuncThatDoesSomethingWithObj(obj *MyObject, number int) bool {
	res, err := obj.actor.DoSomething(number)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return false
	}
	return res
}
