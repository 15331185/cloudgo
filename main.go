package main



import flag "github.com/spf13/pflag"

import "service"

import "os"



const (

	PORT string = "9090"

)



func main() {

	var port string



	// Get predefined env port

	port = os.Getenv("PORT")

	if (len(port) <= 0) {

		port = PORT

	}



	// Define server port

	flag.StringVarP(&port, "port", "p", "9090", "define server port")

	flag.Parse()



	server := service.NewServer()

	server.Run(":" + port)

}
