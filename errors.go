package oddsapi

import (
	"errors"
)

var (
	ErrUnauthenticated = errors.New("Unauthenticated or unauthorized. The API key might be missing or invalid (unauthenticated), or it might at its usage limit (unauthorized). The repsonse body will contain more info")
	ErrQueryParams     = errors.New("One or more of the query params are invalid. The repsonse body will contain more info")
	ErrThrottled       = errors.New("Requests are being sent too frequently - the request was throttled")
	ErrServer          = errors.New("Internal error")
	ErrUnknown         = errors.New("An unrecognized error has occured")
)
