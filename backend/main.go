// Main function for the backend of the operations ecosystem.

package main

import (
	"flag"

	client "capstone.operations_ecosystem/backend/fake_client"
	"capstone.operations_ecosystem/backend/server"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	serverAddrFlag := flag.String("addr", "127.0.0.1", "TCP address for the server to run on.")
	serverPortFlag := flag.Int("port", 9090, "TCP port for the server to run on.")
	serverFlag := flag.Bool("is_server", true, "Is this terminal for the server or the test client?")
	flag.Parse()

	if *serverFlag {
		server.InitServer(serverAddrFlag, serverPortFlag)
	} else {
		client.TestAdminClient(serverAddrFlag, serverPortFlag)
	}
}
