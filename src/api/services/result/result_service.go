package result

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/ejercicio1/src/api/domain/country"
	"github.com/mercadolibre/ejercicio1/src/api/domain/result"
	"github.com/mercadolibre/ejercicio1/src/api/domain/site"
	"github.com/mercadolibre/ejercicio1/src/api/domain/user"
	"github.com/mercadolibre/ejercicio1/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
	"sync"
)


// Get only with Recievers
func GetResponse(userId int64, urlType string) (*result.Result, *apierrors.ApiError){


	user := user.User{
		ID: userId,
	}
	user.Get(urlType)

	site := site.Site{
		ID: user.SiteID,
	}

	country := country.Country{
		ID: user.CountryID,
	}
	site.Get(urlType)
	country.Get(urlType)

	resp := result.Result{
		User: &user,
		Site: &site,
		Country: &country,
	}

	return &resp, nil
}

// Get with WaitGroups
func GetResponseWg(userId int64, urlType string) (*result.Result, *apierrors.ApiError){

	ch := make(chan *result.Result)
	defer close(ch)

	var wg sync.WaitGroup
	var result result.Result

	user := user.User{
		ID: userId,
	}
	user.Get(urlType)

	site := site.Site{
		ID: user.SiteID,
	}

	country := country.Country{
		ID: user.CountryID,
	}

	wg.Add(2)
	go site.GetWg(urlType, &wg)
	go country.GetWg(urlType, &wg)
	wg.Wait()

	result.User = &user
	result.Country = &country
	result.Site = &site

	return &result, nil
}


// Get with WaitGroups and one Channel
func GetResponseCh(userId int64, urlType string) (*result.Result, *apierrors.ApiError){

	ch := make(chan *result.Result)
	defer close(ch)

	var wg sync.WaitGroup
	var result result.Result


	user := user.User{
		ID: userId,
	}
	user.Get(urlType)
	result.User = &user

	go func() {

		for i := 0; i < 2; i++ {

			item := <-ch
			wg.Done()
			if item.Site != nil {
				result.Site = item.Site
				continue
			}

			if item.Country != nil {
				result.Country = item.Country
				continue
			}
		}
	}()

	wg.Add(2)
	go getCountry(user.CountryID, urlType, ch, &wg)
	go getSite(user.SiteID, urlType, ch, &wg)
	wg.Wait()

	return &result, nil

}


/*
func GetResponseCh2(userId int64, urlType string) (*result.Result, *apierrors.ApiError){

	ch := make(chan *result.Result)
	defer close(ch)

	var wg sync.WaitGroup
	var result result.Result

	user := user.User{
		ID: userId,
	}
	user.Get(urlType)

	wg.Add(2)
	go getSite(user.SiteID, urlType, ch, &wg)
	go getCountry(user.CountryID, urlType, ch, &wg)
	wg.Wait()

	return &result, nil
}
*/

func getSite(siteId string, urlType string, ch chan *result.Result, wg *sync.WaitGroup){

	var site site.Site

	urlSite :=  "https://api.mercadolibre.com/sites/"
	if urlType == "mock" {
		urlSite =  "http://localhost:8081/sites/"
		fmt.Println("URL MOCK ----------")
	}

	url := fmt.Sprintf("%s%s", urlSite, siteId)
	resp, err := http.Get(url)
	if err != nil{

	}
	data, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &site)

	ch <- &result.Result{
		Site: &site,
	}
}

func getCountry (countryId string, urlType string, ch chan *result.Result, wg *sync.WaitGroup) {

	var country country.Country

	urlCountry := "https://api.mercadolibre.com/countries/"
	if urlType == "mock" {
		urlCountry = "http://localhost:8081/sites/"
		fmt.Println("URL MOCK ----------")
	}

	url := fmt.Sprintf("%s%s", urlCountry, countryId)

	resp, err := http.Get(url)
	if err != nil {
		ch <- nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &country)

	ch <- &result.Result{
		Country: &country,
	}
}