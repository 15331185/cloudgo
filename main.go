package main
import flag "github.com/spf13/pflag"
import "github.com/15331185/cloudgo/service"
import "os"

const (
	PORT string = "9090"
)

func main() {
	var port string
	// 获取 env port
	port = os.Getenv("PORT")
	if (len(port) <= 0) {
		port = PORT
	}
	//服务器端口为9090
	flag.StringVarP(&port, "port", "p", "9090", "define server port")
	flag.Parse()
	server := service.NewServer()
	server.Run(":" + port)
}
