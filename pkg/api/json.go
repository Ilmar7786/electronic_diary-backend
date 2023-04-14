package api

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ParseErrorType = map[string]string

func ParseAndValidateJSON[T validation.Validatable](c *gin.Context) (T, error) {
	var input T
	if err := c.ShouldBindJSON(&input); err != nil {
		return input, err
	}

	if err := input.Validate(); err != nil {
		return input, err
	}

	return input, nil
}

type ResponseError struct {
	Errors []ParseErrorType `json:"errors"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ResponseError{Errors: ParseError(message)})
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
