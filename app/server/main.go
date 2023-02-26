package main

import (
	yadis "github.com/liuf66/yadis/server"
)

func main() {
	server := yadis.Server{}
	server.Serve()
}
