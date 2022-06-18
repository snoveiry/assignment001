
package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader   = "Authorization"
	BearerTokenContextKey = "token"
)


func (s *Service) UniversalAuthenticator(c *gin.Context) *string {
	var bearerToken string
	authHeader := c.GetHeader(authorizationHeader)
	if authHeader != "" && strings.Contains(strings.ToLower(authHeader), "bearer") {
		authBearer := strings.SplitN(authHeader, " ", 2)
		if len(authBearer) >= 2 {
			bearerToken = authBearer[1]
		}
	}

	return &bearerToken
}

