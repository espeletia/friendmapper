package domain

import "net/http"

type Error struct {
	Message string
	Code    int
}

func (e Error) Error() string { return e.Message }

var (
	InvalidPassword    = Error{Message: "Password can not be validated", Code: http.StatusNotAcceptable}
	InvalidCredentials = Error{Message: "Invalid credentials", Code: http.StatusUnauthorized}
	NotActivatedUser   = Error{Message: "Please activate your account", Code: http.StatusUnauthorized}
	BadRequest         = Error{Message: "Bad request", Code: http.StatusBadRequest}
	ErrInvalidEmail    = Error{Message: "Invalid email", Code: http.StatusNotAcceptable}
	Unauthorized       = Error{Message: "Unauthorized", Code: http.StatusUnauthorized}
	Forbidden          = Error{Message: "Forbidden", Code: http.StatusForbidden}

	UserNotFound            = Error{Message: "User not found", Code: http.StatusNotFound}
	ErrUserAlreadyVerified  = Error{Message: "User already verified", Code: http.StatusNotAcceptable}
	ErrUserNotEnoughBalance = Error{Message: "User does not have enough balance", Code: http.StatusNotAcceptable}

	UserGroupNotFound = Error{Message: "User group not found", Code: http.StatusNotFound}

	OrderNotActivated = Error{Message: "Unable to activate order", Code: http.StatusNotAcceptable}
	OrderNotDeleted   = Error{Message: "Unable to delete order", Code: http.StatusNotAcceptable}
	OrderNotRefunded  = Error{Message: "Unable to refund order please contact support", Code: http.StatusNotAcceptable}
	OrderInvalidDate  = Error{Message: "Unable to create an order with this date", Code: http.StatusNotAcceptable}
	OrdersExist       = Error{Message: "Unable to delete menu because orders exist", Code: http.StatusBadRequest}

	MenuNotFound = Error{Message: "Menu not found", Code: http.StatusNotFound}

	NoMailFound   = Error{Message: "No mail to retrieve", Code: http.StatusNotFound}
	NoMailPayload = Error{Message: "No mail Payload", Code: http.StatusNotFound}

	ErrorCreatingPriceLvlGroup = Error{Message: "Error creating price level group", Code: http.StatusNotAcceptable}
	ErrPriceLvlGroupNotFound   = Error{Message: "Price level group not found", Code: http.StatusNotFound}

	ErrProtectedGroup    = Error{Message: "Protected group can not be deleted", Code: http.StatusNotAcceptable}
	ErrUserGroupNotFound = Error{Message: "User group not found", Code: http.StatusNotFound}

	ErrProtectedFoodPriceLvl = Error{Message: "Protected food price level can not be deleted", Code: http.StatusNotAcceptable}

	ErrProtectedPriceLvlGroup = Error{Message: "Protected price level group can not be deleted", Code: http.StatusNotAcceptable}

	ErrNoVersionsInDB = Error{Message: "No versions in database", Code: http.StatusNotFound}

	ErrFioApi = Error{Message: "Error while calling Fio API", Code: http.StatusNotFound}
)
