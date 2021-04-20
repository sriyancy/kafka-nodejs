package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "HEAD"}
	switch r.Method {
	case "GET":
		println("GET request")
		fallthrough
	case "DELETE":
		println("DELETE  request")
	case "POST":
		println("POST request")
	case "PUT":
		println("PUT request")
	default:
		println("unhandled method")
	}
}
