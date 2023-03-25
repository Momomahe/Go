package main

import "challenge6/routers"

func main() {
	var PORT = ":4000"
	routers.StartServer().Run(PORT)
}
