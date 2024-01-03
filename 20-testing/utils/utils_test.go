package utils

import (
	"testing"
)

/*
func Test_IsPrime_71(t *testing.T) {
	// Arrange
	no := 71
	expectedResult := true

	//Act
	actualResult := IsPrime(no)

	//Assert
	if actualResult != expectedResult {

			// t.Logf("IsPrime(%d) = %t, but got %t\n", no, expectedResult, actualResult)
			// t.Fail()

		t.Errorf("IsPrime(%d) = %t, but got %t\n", no, expectedResult, actualResult)
	}
}

func Test_IsPrime_97(t *testing.T) {
	// Arrange
	no := 97
	expectedResult := true

	//Act
	actualResult := IsPrime(no)

	//Assert
	if actualResult != expectedResult {

			// t.Logf("IsPrime(%d) = %t, but got %t\n", no, expectedResult, actualResult)
			// t.Fail()

		t.Errorf("IsPrime(%d) = %t, but got %t\n", no, expectedResult, actualResult)
	}
}
*/

func Test_IsPrime(t *testing.T) {
	testData := []struct {
		name           string
		no             int
		actualResult   bool
		expectedResult bool
	}{
		{name: "Testing_IsPrime_71", no: 71, expectedResult: true},
		{name: "Testing_IsPrime_90", no: 90, expectedResult: false},
		{name: "Testing_IsPrime_13", no: 13, expectedResult: false},
		{name: "Testing_IsPrime_91", no: 91, expectedResult: false},
	}

	for _, td := range testData {
		t.Run(td.name, func(t *testing.T) {
			td.actualResult = IsPrime(td.no)

			//Assert
			if td.actualResult != td.expectedResult {
				t.Errorf("IsPrime(%d) = %t, but got %t\n", td.no, td.expectedResult, td.actualResult)
			}
		})
	}
}

// Benchmarking
func Benchmark_IsPrime_97(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(97)
	}
}

func Benchmark_IsPrime_2_97(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime_2(97)
	}
}
