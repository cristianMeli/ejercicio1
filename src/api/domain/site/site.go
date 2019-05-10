package site

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/ejercicio1/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
	"sync"
)

type Site struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	CountryID          string   `json:"country_id"`
	SaleFeesMode       string   `json:"sale_fees_mode"`
	MercadopagoVersion int      `json:"mercadopago_version"`
	DefaultCurrencyID  string   `json:"default_currency_id"`
	ImmediatePayment   string   `json:"immediate_payment"`
	PaymentMethodIds   []string `json:"payment_method_ids"`
	Settings           struct {
		IdentificationTypes      []string      `json:"identification_types"`
		TaxpayerTypes            []interface{} `json:"taxpayer_types"`
		IdentificationTypesRules interface{}   `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Categories []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"categories"`
}

func (site *Site) GetWg(urlType string, wg *sync.WaitGroup) *apierrors.ApiError{

	defer wg.Done()
	if site.ID == ""{
		return &apierrors.ApiError{
			Message: "Country Id is empty",
			Status: http.StatusBadRequest,
		}
	}

	urlSite := "http://api.mercadolibre.com/sites/"
	if urlType == "mock" {
		urlSite = "http://localhost:8081/sites/"
		fmt.Println("URL MOCK ----------")
	}
	url := fmt.Sprintf("%s%s", urlSite, site.ID)

	response, err := http.Get(url)
	if err != nil {
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

	if err = json.Unmarshal(data, &site); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}

func (site *Site) Get(urlType string) *apierrors.ApiError{

	if site.ID == ""{
		return &apierrors.ApiError{
			Message: "Country Id is empty",
			Status: http.StatusBadRequest,
		}
	}

	urlSite := "http://api.mercadolibre.com/sites/"
	if urlType == "mock" {
		urlSite = "http://localhost:8081/sites/"
		fmt.Println("URL MOCK ----------")
	}
	url := fmt.Sprintf("%s%s", urlSite, site.ID)

	response, err := http.Get(url)
	if err != nil {
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

	if err = json.Unmarshal(data, &site); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return nil
}