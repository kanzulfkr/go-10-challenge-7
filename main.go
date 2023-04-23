package main

import "C7/routers"

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
