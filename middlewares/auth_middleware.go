package middlewares

import (
	"net/http"
	"goauth/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Obtém o token do cookie
		tokenString, err := ctx.Cookie("token")
		if err != nil {
			// Se não encontrar o cookie ou houver erro ao acessá-lo, retorna erro
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// Valida o token com sua função de validação
		_, err = utils.ValidateToken(tokenString)
		if err != nil {
			// Se o token for inválido, retorna erro
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// Se o token for válido, continua para o próximo handler
		ctx.Next()
	}
}
