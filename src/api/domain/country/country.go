package country

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/ejercicio1/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
	"sync"
)

type Country struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Locale             string `json:"locale"`
	CurrencyID         string `json:"currency_id"`
	DecimalSeparator   string `json:"decimal_separator"`
	ThousandsSeparator string `json:"thousands_separator"`
	TimeZone           string `json:"time_zone"`
	GeoInformation     struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
	} `json:"geo_information"`
	States []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"states"`
}

func (country *Country) GetWg(urlType string, wg *sync.WaitGroup) *apierrors.ApiError{

	defer wg.Done()
	if country.ID == "" {
		return &apierrors.ApiError{
			Message: "Country ID is empty",
			Status: http.StatusBadRequest,
		}
	}
	urlCountry := "http://api.mercadolibre.com/countries/"
	if urlType == "mock" {
		urlCountry = "http://localhost:8081/countries/"
		fmt.Println("URL MOCK ----------")
	}
	url := fmt.Sprintf("%s%s", urlCountry, country.ID)

	response, err := http.Get(url)
	if err != nil{

		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {

		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err = json.Unmarshal(data, &country); err != nil{

		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}

func (country *Country) Get(urlType string) *apierrors.ApiError{

	if country.ID == "" {
		return &apierrors.ApiError{
			Message: "Country ID is empty",
			Status: http.StatusBadRequest,
		}
	}

	urlCountry := "http://api.mercadolibre.com/countries/"
	if urlType == "mock" {
		urlCountry = "http://localhost:8081/countries/"
		fmt.Println("URL MOCK ----------")
	}

	url := fmt.Sprintf("%s%s", urlCountry, country.ID)

	response, err := http.Get(url)
	if err != nil{

		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {

		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err = json.Unmarshal(data, &country); err != nil{

		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}