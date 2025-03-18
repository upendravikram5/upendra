Unit Testing in Golang
Unit testing in Golang is done using the built-in testing package. You can write test cases in a file that ends with _test.go and run them using go test.

1. Writing a Simple Unit Test
Let's say you have a simple function in mathutil.go:

go
Copy
Edit
package mathutil

// Add returns the sum of two numbers
func Add(a, b int) int {
    return a + b
}
Now, create a test file named mathutil_test.go:

go
Copy
Edit
package mathutil

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2,3) = %d; want %d", result, expected)
    }
}
2. Running the Test
Run the test using:

sh
Copy
Edit
go test
or

sh
Copy
Edit
go test -v  # For verbose output
3. Using t.Run for Multiple Test Cases
Instead of writing multiple test functions, use t.Run:

go
Copy
Edit
func TestAddCases(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"Both positive", 2, 3, 5},
        {"One negative", -1, 4, 3},
        {"Both negative", -2, -3, -5},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
4. Using assert with github.com/stretchr/testify (Recommended)
You can use testify to simplify assertions:

Install Testify:
sh
Copy
Edit
go get github.com/stretchr/testify
Refactored Test with assert
go
Copy
Edit
package mathutil

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestAddWithAssert(t *testing.T) {
    assert.Equal(t, 5, Add(2, 3), "2 + 3 should be 5")
    assert.Equal(t, -1, Add(-2, 1), "-2 + 1 should be -1")
}
5. Testing for Errors with assert.Panics
If a function should panic on an invalid input:

go
Copy
Edit
func Divide(a, b int) int {
    if b == 0 {
        panic("division by zero")
    }
    return a / b
}

func TestDivide(t *testing.T) {
    assert.Panics(t, func() { Divide(4, 0) }, "Expected panic when dividing by zero")
}
6. Benchmarking in Golang
You can also add benchmarks in the same test file using testing.B:

go
Copy
Edit
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
Run benchmarks using:

sh
Copy
Edit
go test -bench .
7. Mocking Dependencies in Golang
For testing functions that depend on external services (e.g., databases, APIs), use mocking with github.com/stretchr/testify/mock.

Example:
go
Copy
Edit
package service

import (
    "testing"

    "github.com/stretchr/testify/mock"
)

// Define a mock struct
type MockDB struct {
    mock.Mock
}

func (m *MockDB) GetUser(id int) string {
    args := m.Called(id)
    return args.String(0)
}

func TestGetUser(t *testing.T) {
    mockDB := new(MockDB)
    mockDB.On("GetUser", 1).Return("Alice")

    result := mockDB.GetUser(1)
    assert.Equal(t, "Alice", result, "Expected Alice")
}
Final Thoughts
Use testing.T for basic tests.
Use github.com/stretchr/testify/assert for cleaner assertions.
Use testing.B for benchmarks.
Use github.com/stretchr/testify/mock for mocking dependencies.
