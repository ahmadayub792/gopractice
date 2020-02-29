package weather_test

import (
	"testing"

	"github.com/ahmadayub792/gopractice/weather"
	"github.com/stretchr/testify/assert"
)

type MockAPI struct {
	B   []byte
	Err error
}

func (o MockAPI) GET(url string) ([]byte, error) {
	return o.B, o.Err
}

type TestCase struct {
	Input         weather.Weather
	ExpectedTemp  int
	ExpectedError error
}

func TestWeather(t *testing.T) {
	testcases := []TestCase{
		{
			Input: weather.Weather{
				Request: MockAPI{B: []byte(`123`), Err: nil},
			},
			ExpectedTemp:  246,
			ExpectedError: nil,
		}, {
			Input: weather.Weather{
				Request: MockAPI{B: []byte(`0`), Err: nil},
			},
			ExpectedTemp:  0,
			ExpectedError: nil,
		}, {
			Input: weather.Weather{
				Request: MockAPI{B: []byte(`-10`), Err: nil},
			},
			ExpectedTemp:  -20,
			ExpectedError: nil,
		},
	}

	for _, testcase := range testcases {
		temp, err := testcase.Input.Temperature("Karachi")
		assert.Equal(t, testcase.ExpectedTemp, temp, "Temp did not match")
		assert.Equal(t, testcase.ExpectedError, err, "Error did not match")

	}
}
