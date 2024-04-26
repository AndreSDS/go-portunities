package handler

import (
	"fmt"

	models "github.com/andresds/go-portunities/handler/models"
)

func Validate(request *models.OpenningRequest) error {
	if request == nil {
		return fmt.Errorf("malformad request body")
	}

	if request.Role == "" {
		return ErrParamsIsRequired("role", "string")
	}
	if request.Company == "" {
		return ErrParamsIsRequired("company", "string")
	}
	if request.Location == "" {
		return ErrParamsIsRequired("location", "string")
	}
	if request.Link == "" {
		return ErrParamsIsRequired("link", "string")
	}
	if request.Remote == nil {
		return ErrParamsIsRequired("remote", "bool")
	}

	// cheack if field salary exists
	if request.Salary <= 0 {
		return ErrParamsIsRequired("salary", "int")
	}
	return nil
}
