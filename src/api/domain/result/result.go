package result

import(
	"github.com/mercadolibre/ejercicio1/src/api/domain/country"
	"github.com/mercadolibre/ejercicio1/src/api/domain/site"
	"github.com/mercadolibre/ejercicio1/src/api/domain/user"

)

type Result struct{
	User *user.User
	Country *country.Country
	Site *site.Site
}
