package dto

import (
	"net/http"
	"strconv"
	"strings"
)

// It's a representation query string params to filter and paginate products
type PaginationRequestParams struct {
    Search       string   `json:"search"`
    Descending   []string `json:"descending"`
    Page         int      `json:"page"`
    ItemsPerPage int      `json:"itemsPerPage"`
    Sort         []string `json:"sort"`
}

// Converts query string params to a PaginationRequestParams struct
func FromValuePaginationRequestParams(request *http.Request) (*PaginationRequestParams, error) {
	// Get values from query string params
	page, _ := strconv.Atoi(request.FormValue("page"))
    itemsPerPage, _ := strconv.Atoi(request.FormValue("itemsPerPage"))

	// Create a PaginationRequestParms struct
    paginationRequestParms := PaginationRequestParams{
        Search:       request.FormValue("search"),
        Descending:   strings.Split(request.FormValue("descending"), ","),
        Sort:         strings.Split(request.FormValue("sort"), ","),
        Page:         page,
        ItemsPerPage: itemsPerPage,
    }

	// Return the PaginationRequestParms struct
    return &paginationRequestParms, nil
}