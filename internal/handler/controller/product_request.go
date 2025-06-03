package controller

import "strconv"

type GetProductsRequest struct {
	Page  int
	Limit int
}

func ValidateGetProductsRequest(page string, limit string) *GetProductsRequest {
	const (
		defaultPage  = 1
		defaultLimit = 10
		minPage      = 1
		minLimit     = 1
		maxLimit     = 100
	)

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < minPage {
		pageInt = defaultPage
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < minLimit || limitInt > maxLimit {
		limitInt = defaultLimit
	}

	return &GetProductsRequest{
		Page:  pageInt,
		Limit: limitInt,
	}
}
