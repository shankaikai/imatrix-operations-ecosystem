// Main function for the backend of the operations ecosystem.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	client "capstone.operations_ecosystem/backend/fake_client"
	"capstone.operations_ecosystem/backend/fake_server"
	"capstone.operations_ecosystem/backend/server"
	tclient "capstone.operations_ecosystem/backend/telegram_client_tests"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	serverAddrFlag := flag.String("addr", "0.0.0.0", "TCP address for the server to run on.")
	serverPortFlag := flag.Int("port", 9090, "TCP port for the server to run on.")
	teleServerAddrFlag := flag.String("tele_addr", "192.168.100.26", "TCP address for the server to run on.")
	teleServerPortFlag := flag.Int("tele_port", 9091, "TCP port for the server to run on.")
	webProxyServerAddrFlag := flag.String("wproxy_addr", "0.0.0.0", "TCP address for the web proxy server to run on.")
	webProxyServerPortFlag := flag.Int("wproxy_port", 9089, "TCP port for the web proxy server to run on.")
	testLEDAddrFlag := flag.String("led_addr", "http://192.168.100.123", "IP address of the LED lights on the AIFS for testing")
	serverFlag := flag.Bool("is_server", true, "Is this terminal for the server or the test client?")
	fakeServerFlag := flag.Bool("is_fserver", false, "Start the server or the fake server?")
	teleClientFlag := flag.Bool("is_tclient", false, "Should we test the server as a telebot client?")
	cliFlag := flag.Bool("use_cli", false, "Should we start a CLI?")
	startProxyFlag := flag.Bool("is_proxy", false, "Should we start the proxy or the main backend server?")

	flag.Parse()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	if *fakeServerFlag {
		if *cliFlag {
			go fake_server.InitServer(serverAddrFlag, serverPortFlag)
		} else {
			fake_server.InitServer(serverAddrFlag, serverPortFlag)
		}
	} else if *teleClientFlag {
		tclient.TestTelegramBroadcasts(teleServerAddrFlag, teleServerPortFlag)
		// tclient.TestTelegramRosters(teleServerAddrFlag, teleServerPortFlag)
	} else if *startProxyFlag{
		if *cliFlag {
			go server.Proxy_main(serverAddrFlag, serverPortFlag, webProxyServerAddrFlag, webProxyServerPortFlag)
		} else {
			server.Proxy_main(serverAddrFlag, serverPortFlag, webProxyServerAddrFlag, webProxyServerPortFlag)
		}
	} else if *serverFlag {
		if *cliFlag {
			go server.InitServer(serverAddrFlag, serverPortFlag, teleServerAddrFlag, teleServerPortFlag, testLEDAddrFlag, webProxyServerAddrFlag, webProxyServerPortFlag)
		} else {
			server.InitServer(serverAddrFlag, serverPortFlag, teleServerAddrFlag, teleServerPortFlag, testLEDAddrFlag, webProxyServerAddrFlag, webProxyServerPortFlag)
		}
	} else {
		// client.TestAdminClientUser(serverAddrFlag, serverPortFlag)
		// client.TestAdminClientClient(serverAddrFlag, serverPortFlag)
		client.TestBroadcastClient(serverAddrFlag, serverPortFlag)
		// client.TestRosteringClient(serverAddrFlag, serverPortFlag)
		// client.TestIncidentReportClient(serverAddrFlag, serverPortFlag)
		client.TestCameraIotClientUser(serverAddrFlag, serverPortFlag)
	}

	//The only reason these are here is to help ensure a clean exit (no zombie processes)
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Terminating...")
		os.Exit(0)
	}()

	reader := bufio.NewReader(os.Stdin)
mainLoop:
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if len(text) == 0 {
			continue
		}
		tokenised := strings.Split(text, " ")
		cmd := tokenised[0]
		switch cmd {
		case "exit":
			fmt.Println("Terminating...")
			break mainLoop
		default:
			fmt.Println("Unrecognised command.")
		}
	}
}
