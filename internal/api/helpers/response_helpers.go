package helpers

import (
	"net/http"

	"github.com/LamichhaneBibek/go-template/internal/constants"
)

type BaseHttpResponse struct {
	Result     any                  `json:"result"`
	Success    bool                 `json:"success"`
	ResultCode constants.ResultCode `json:"resultCode"`
	// ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Error any `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode constants.ResultCode) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode constants.ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}

}

func GenerateBaseResponseWithAnyError(result any, success bool, resultCode constants.ResultCode, err any) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode constants.ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
		// ValidationErrors: validations.GetValidationErrors(err),
	}
}

var StatusCodeMapping = map[string]int{
	// User
	constants.EmailExists:      409,
	constants.UsernameExists:   409,
	constants.RecordNotFound:   404,
	constants.PermissionDenied: 403,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
