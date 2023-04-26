package main
import(
	"fmt"
	"github.com/Karthika-Rajagopal/jwt-go/initializers"

)

func init(){
	initializers.LoadEnvVariables()
}
func main(){
	fmt.Println("Hello")
}