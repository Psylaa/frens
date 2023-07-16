package models

import "github.com/gofiber/fiber/v2"

type ErrorDetail string

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// Static Error Definitions
var (
	// General Errors
	ErrBadRequest = &Error{
		Status: fiber.StatusBadRequest,
		Title:  "Bad Request",
		Detail: "Your request could not be understood or was missing required parameters.",
	}
	ErrUnauthorized = &Error{
		Status: fiber.StatusUnauthorized,
		Title:  "Unauthorized",
		Detail: "Authentication failed or user doesn't have permissions for the request.",
	}
	ErrForbidden = &Error{
		Status: fiber.StatusForbidden,
		Title:  "Forbidden",
		Detail: "Authentication succeeded but authenticated user doesn't have access to the resource.",
	}
	ErrNotFound = &Error{
		Status: fiber.StatusNotFound,
		Title:  "Not Found",
		Detail: "A specified resource could not be found.",
	}
	ErrInternalServerError = &Error{
		Status: fiber.StatusInternalServerError,
		Title:  "Internal Server Error",
		Detail: "There was an error processing the request.",
	}

	// Validation errors
	ErrInvalidBody = &Error{
		Status: fiber.StatusBadRequest,
		Title:  "Invalid Body",
		Detail: "Unable to parse the request body. Please check the body and try again.",
	}
	ErrInvalidCredentials = &Error{
		Status: fiber.StatusBadRequest,
		Title:  "Invalid Credentials",
		Detail: "The provided credentials are incorrect. Please try again.",
	}
	ErrInvalidEmailFormat = &Error{
		Status: fiber.StatusBadRequest,
		Title:  "Invalid Email Format",
		Detail: "The email provided is not in a valid format. Please correct it and try again.",
	}
	ErrInvalidPhoneFormat = &Error{
		Status: fiber.StatusBadRequest,
		Title:  "Invalid Phone Format",
		Detail: "The phone number provided is not in a valid format. Please correct it and try again.",
	}
	ErrInvalidPasswordLength = &Error{
		Status: fiber.StatusBadRequest,
		Title:  "Invalid Password Length",
		Detail: "The password provided is not long enough. It should be at least 8 characters.",
	}
	ErrInvalidPostID = &Error{
		Status: fiber.StatusBadRequest,
		Title:  "Invalid Post ID",
		Detail: "The post ID provided is not valid. Please check the ID and try again.",
	}

	// Authentication and Authorization errors
	ErrInvalidToken = &Error{
		Status: fiber.StatusUnauthorized,
		Title:  "Invalid Token",
		Detail: "The token provided is not valid. Please reauthenticate and try again.",
	}
	ErrUserNotActivated = &Error{
		Status: fiber.StatusUnauthorized,
		Title:  "User Not Activated",
		Detail: "This account has not been activated yet. Please activate your account and try again.",
	}
	ErrUserSuspended = &Error{
		Status: fiber.StatusForbidden,
		Title:  "User Suspended",
		Detail: "This account has been suspended. Please contact support for further assistance.",
	}
	ErrTokenExpired = &Error{
		Status: fiber.StatusUnauthorized,
		Title:  "Token Expired",
		Detail: "The token provided has expired. Please reauthenticate and try again.",
	}
	ErrTokenInvalid = &Error{
		Status: fiber.StatusUnauthorized,
		Title:  "Invalid Token",
		Detail: "The token provided is not valid. Please reauthenticate and try again.",
	}

	// User errors
	ErrUserExists = &Error{
		Status: fiber.StatusConflict,
		Title:  "User Already Exists",
		Detail: "A user with this email already exists. Please use a different email address.",
	}
	ErrUserNotFound = &Error{
		Status: fiber.StatusNotFound,
		Title:  "User Not Found",
		Detail: "We couldn't find a user with the provided details. Please check the details and try again.",
	}
)

// Generates and sends a JSON API error response based on provided error
// and optional detail message. If no detail is provided, generic detail
// message will be used.
func (e *Error) SendResponse(c *fiber.Ctx, customDetail ...string) error {
	if len(customDetail) > 0 {
		e.Detail = customDetail[0]
	}
	return c.Status(e.Status).JSON(ErrorResponse{Errors: []Error{*e}})
}
