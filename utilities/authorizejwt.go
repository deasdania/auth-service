package utilities

import (
	"account-metalit/api/auth/usecase"
	"account-metalit/try"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func TokenAuthMiddleware(authUsecase usecase.IAuthUsecase) gin.HandlerFunc {
	fmt.Println("TOKENAUTH")
	return func(c *gin.Context) {
		var err error
		try.This(func() {
			metadata, errA := authUsecase.ExtractTokenMetadata(c.Request)
			if errA != nil {
				err = errA
				return
			}
			_, errA = authUsecase.FetchAuth(metadata)
			if errA != nil {
				err = errA
				return
			}
		}).Catch(func(errE try.E) {
			err = errors.New("error")
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
			c.Abort()
		}
		c.Next()

	}
}

func CheckRestClientJWT(authUsecase usecase.IAuthUsecase) gin.HandlerFunc {
	fmt.Println("CheckRestClient")
	return func(c *gin.Context) {
		mekarClient := c.Request.Header.Get("Authorization")
		if mekarClient != "" {
			// fmt.Println(mekarClient, "METALIT_CLIENT")
			var err error
			try.This(func() {
				metadata, errA := authUsecase.ExtractTokenMetadata(c.Request)
				if errA != nil {
					err = errA
					return
				}
				_, errA = authUsecase.FetchAuth(metadata)
				if errA != nil {
					err = errA
					return
				}
			}).Catch(func(errE try.E) {
				err = errors.New("error")
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
				c.Abort()
			}
			c.Next()
		} else {
			user, password, ok := c.Request.BasicAuth()
			if user == os.Getenv("ACCOUNT_SERVICE_USERNAME") && password == os.Getenv("ACCOUNT_SERVICE_PASSWORD") {
				c.Next()
			}
			if !ok {
				c.JSON(http.StatusUnauthorized, "You need to be authorized to access this route")
				c.Abort()
			}
		}
	}
}