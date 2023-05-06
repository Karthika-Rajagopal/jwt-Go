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
	tokenString, err := c.Cookie("Authorization")  //Reads the Authorization cookie from the incoming HTTP request using c.Cookie("Authorization") method
	if err !=nil{
		c.AbortWithStatus(http.StatusUnauthorized)  //If the cookie is not present, aborts the request with HTTP status code 401 Unauthorized using c.AbortWithStatus(http.StatusUnauthorized)
	}
	//Decode/Validate it
	// Parse takes the token string and a function for looking up the key. 
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {  //If the cookie is present, it decodes and validates it using jwt.Parse() method, passing in the token string and a function that returns the secret key used for decoding
	
	
	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
 	return []byte(os.Getenv("SECRET")), nil
 })

if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid { //If the token is valid and claims can be extracted, the ok variable will be set to true and the claims variable will hold the claims extracted from the token.
	
	//check the exp
	if float64(time.Now().Unix()) > claims["exp"].(float64){
		c.AbortWithStatus(http.StatusUnauthorized)
}
	}
	//Find the user with token sub
	var user models.User
	initializers.DB.First(&user,claims["sub"])  // it retrieves the user associated with the JWT token by querying the database for the user with the sub claim in the token
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



