package main

import( "net/http"
	"day2/controllers"	
)

// "errors"

func main() {
	// u := models.User{
	// 	ID:        2,
	// 	FirstName: "tricia",
	// 	LastName:  "MC",
	// }
	// fmt.Println(u)
	// port := 3000
	// _, err := startWebServer(port, 2)
	// fmt.Println(err)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

//creating func
// func startWebServer(port, numberOfRetries int) (int, error) {
// 	fmt.Println("Starting server...")
// 	fmt.Println("Server started ", port)
// 	fmt.Println("retries ", numberOfRetries)
// 	// return errors.New("Someting went wrong")
// 	return port, nil
// }
