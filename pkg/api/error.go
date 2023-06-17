package api

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	ResponseErrors struct {
		Errors []ParseErrorType `json:"errors"`
	}

	ResponseError struct {
		Error string `json:"error"`
	}

	ParseErrorType = map[string]string

	Response struct {
		Message string `json:"message"`
	}
)

func NewErrorsResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, ResponseErrors{Errors: ParseError(message)})
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, ResponseError{Error: message})
}

func NewError(str string) ResponseError {
	return ResponseError{Error: str}
}

func ParseError(str string) []ParseErrorType {
	if !strings.Contains(str, ":") {
		return []ParseErrorType{{"fail": str}}
	}

	items := strings.Split(str, "; ")

	var result []ParseErrorType

	for _, item := range items {
		re := regexp.MustCompile(`^(.*?):\s+(.*)$`)
		match := re.FindStringSubmatch(item)

		if len(match) == 3 {
			result = append(result, map[string]string{match[1]: match[2]})
		}
	}

	return result
}
