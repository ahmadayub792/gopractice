package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type GetRequestMaker interface {
	GET(string) ([]byte, error)
}

type OpenWeatherAPI struct {
	APIKey string
}

func (o *OpenWeatherAPI) GET(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

type Weather struct {
	Request GetRequestMaker
}

func (w *Weather) Temperature(city string) (int, error) {
	buf, err := w.Request.GET(fmt.Sprintf("http://google.com?q=%s", city))
	if err != nil {
		return 0, err
	}

	temp, err := strconv.Atoi(string(buf))
	if err != nil {
		return 0, err
	}

	return temp * 2, nil
}

func abc () {
	weather{
		Request: OpenWeatherAPI{APIKey: }
	}
}