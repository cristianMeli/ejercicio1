package user

import(
	"github.com/mercadolibre/ejercicio1/src/api/domain/user"
	"github.com/mercadolibre/ejercicio1/src/api/utils/apierrors"
)

const urlUsers = "https://api.mercadolibre.com/users/"

func GetUser(userId int64) (*user.User, *apierrors.ApiError){

	user := user.User{
		ID: userId,
	}
	if apiErr := user.Get(); apiErr != nil {
		return nil, apiErr
	}

	return &user, nil
}
