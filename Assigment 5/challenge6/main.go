package main

import "challenge5/routers"

func main() {
	var PORT = ":4000"
	routers.StartServer().Run(PORT)
}
