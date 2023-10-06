package domain

// It's representation of Fetch methods returns
type Pagination[T any] struct {
    Items T     `json:"items"`
    Total int32 `json:"total"`
}