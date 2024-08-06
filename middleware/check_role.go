package middleware

import (
	"context"
	"net/http"

	"github.com/anhhuy1010/cms-menu/grpc"
	pbUsers "github.com/anhhuy1010/cms-user/grpc/proto/users"
	"github.com/gin-gonic/gin"
)

func CheckRole(token string) (*pbUsers.DetailResponse, error) {
	grpcConn := grpc.GetInstance()
	client := pbUsers.NewUserClient(grpcConn.UserConnect)
	req := pbUsers.DetailRequest{
		Token: token,
	}
	resp, err := client.Detail(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func RoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("x-token")
		if token == "" {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		resp, err := CheckRole(token)
		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		if resp == nil {

			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token response"})
			c.Abort()
			return
		}
		c.Set("userRole", resp.Role)
		c.Set("userUuid", resp.UserUuid)

		if resp.Role != "admin" && c.Request.Method != http.MethodGet {

			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}
