package middlewares

import (
	"Back-end/config"
	"Back-end/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// VerifyJWT 函数用于验证 JWT 的合法性
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	// 解析 JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回用于验证签名的密钥
		return []byte(config.Config.GetString("jwt.secret")), nil
	})

	// 检查解析过程中是否出现错误
	if err != nil {
		utils.LogError(err)
		return nil, err
	}

	// 检查 JWT 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token is valid")
		fmt.Println("Username:", claims["username"])
		fmt.Println("Type:", int(claims["type"].(float64))) // 注意这里需要将float64转换为int
	} else {
		fmt.Println("Invalid token")
	}
	utils.LogError(err)
	return token, err
}

// TokenAuthMiddleware 是一个中间件函数，用于验证 JWT 的合法性
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			// 如果 Authorization 字段为空，返回 401 Unauthorized 错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// 验证 JWT 的合法性
		token, err := VerifyJWT(tokenString)
		if err != nil || !token.Valid {
			// 如果 JWT 验证失败，返回 401 Unauthorized 错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 检查 JWT 的 claims 是否有效
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			// 如果 JWT 的 claims 无效，返回 401 Unauthorized 错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// 将解析得到的 username 和 type 参数传递给下一步命令
		c.Set("username", claims["username"].(string))
		c.Set("type", int(claims["type"].(float64)))

		// 继续处理请求
		c.Next()
	}
}
