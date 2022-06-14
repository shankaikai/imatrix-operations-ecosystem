// Main function for the backend of the operations ecosystem.

package main

import (
	"flag"

	client "capstone.operations_ecosystem/backend/fake_client"
	"capstone.operations_ecosystem/backend/fake_server"
	"capstone.operations_ecosystem/backend/server"
	tclient "capstone.operations_ecosystem/backend/telegram_client_tests"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	serverAddrFlag := flag.String("addr", "127.0.0.1", "TCP address for the server to run on.")
	serverPortFlag := flag.Int("port", 9090, "TCP port for the server to run on.")
	teleServerAddrFlag := flag.String("tele_addr", "127.0.0.1", "TCP address for the server to run on.")
	teleServerPortFlag := flag.Int("tele_port", 9091, "TCP port for the server to run on.")
	testLEDAddrFlag := flag.String("led_addr", "http://192.168.0.163", "IP address of the LED lights on the AIFS for testing")
	serverFlag := flag.Bool("is_server", true, "Is this terminal for the server or the test client?")
	fakeServerFlag := flag.Bool("is_fserver", false, "Is this terminal for the server or the test client?")
	teleClientFlag := flag.Bool("is_tclient", false, "Is this terminal for the server or the test client?")
	flag.Parse()

	if *fakeServerFlag {
		fake_server.InitServer(serverAddrFlag, serverPortFlag)
	} else if *teleClientFlag {
		// tclient.TestTelegramBroadcasts(teleServerAddrFlag, teleServerPortFlag)
		tclient.TestTelegramRosters(teleServerAddrFlag, teleServerPortFlag)
	} else if *serverFlag {
		server.InitServer(serverAddrFlag, serverPortFlag, teleServerAddrFlag, teleServerPortFlag, testLEDAddrFlag)
	} else {
		// client.TestAdminClientUser(serverAddrFlag, serverPortFlag)
		// client.TestAdminClientClient(serverAddrFlag, serverPortFlag)
		// client.TestBroadcastClient(serverAddrFlag, serverPortFlag)
		client.TestRosteringClient(serverAddrFlag, serverPortFlag)
	}
}
