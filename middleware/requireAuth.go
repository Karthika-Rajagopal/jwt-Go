package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	//Get the cookie of req
	tokenString, err := c.Cookie("Authorization")
	if err !=nil{
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Decode/Validate it
	// Parse takes the token string and a function for looking up the key. 
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	
	
	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
 	return []byte(os.Getenv("SECRET")), nil
 })

if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	
	//check the exp
	if float64(time.Now().Unix()) > claims["exp"].(float64){
		c.AbortWithStatus(http.StatusUnauthorized)
}
	}
	//Find the user with token sub
	var user models.User
	initializers.DB.First(&user,claims["sub"])
	if user.ID ==0 {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Attach to req
	c.Set("user", user)
	//continue

	c.Next()
	
} else {
	c.AbortWithStatus(http.StatusUnauthorized)
}



