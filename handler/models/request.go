package handler

import "fmt"

func ErrParamsIsRequired(
	name, typ string,
) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type OpenningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *OpenningRequest) Validate() error {
	if r == nil {
		return fmt.Errorf("malformad request body")
	}

	if r.Role == "" {
		return ErrParamsIsRequired("role", "string")
	}
	if r.Company == "" {
		return ErrParamsIsRequired("company", "string")
	}
	if r.Location == "" {
		return ErrParamsIsRequired("location", "string")
	}
	if r.Link == "" {
		return ErrParamsIsRequired("link", "string")
	}
	if r.Remote == nil {
		return ErrParamsIsRequired("remote", "bool")
	}

	// cheack if field salary exists
	if r.Salary <= 0 {
		return ErrParamsIsRequired("salary", "int")
	}
	return nil
}
