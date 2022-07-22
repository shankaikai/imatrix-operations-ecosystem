// Main function for the backend of the operations ecosystem.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"

	client "capstone.operations_ecosystem/backend/fake_client"
	"capstone.operations_ecosystem/backend/fake_server"
	"capstone.operations_ecosystem/backend/server"
	tclient "capstone.operations_ecosystem/backend/telegram_client_tests"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	serverAddrFlag := flag.String("addr", "0.0.0.0", "TCP address for the server to run on.")
	serverPortFlag := flag.Int("port", 9090, "TCP port for the server to run on.")
	teleServerAddrFlag := flag.String("tele_addr", "telebot", "TCP address for the server to run on.")
	teleServerPortFlag := flag.Int("tele_port", 9091, "TCP port for the server to run on.")
	webProxyServerAddrFlag := flag.String("wproxy_addr", "0.0.0.0", "TCP address for the web proxy server to run on.")
	webProxyServerPortFlag := flag.Int("wproxy_port", 9089, "TCP port for the web proxy server to run on.")
	testLEDAddrFlag := flag.String("led_addr", "http://192.168.166.238", "IP address of the LED lights on the AIFS for testing")
	serverFlag := flag.Bool("is_server", true, "Is this terminal for the server or the test client?")
	fakeServerFlag := flag.Bool("is_fserver", false, "Start the server or the fake server?")
	teleClientFlag := flag.Bool("is_tclient", false, "Should we test the server as a telebot client?")

	flag.Parse()

	// Set up sentry
	initSentry()
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)
	defer sentry.Recover()

	if *fakeServerFlag {
		go fake_server.InitServer(serverAddrFlag, serverPortFlag)
	} else if *teleClientFlag {
		tclient.TestTelegramBroadcasts(teleServerAddrFlag, teleServerPortFlag)
		// tclient.TestTelegramRosters(teleServerAddrFlag, teleServerPortFlag)
	} else if *serverFlag {
		go server.InitServer(serverAddrFlag, serverPortFlag, teleServerAddrFlag, teleServerPortFlag, testLEDAddrFlag, webProxyServerAddrFlag, webProxyServerPortFlag)
	} else {
		client.TestAdminClientUser(serverAddrFlag, serverPortFlag)
		// client.TestAdminClientClient(serverAddrFlag, serverPortFlag)
		// client.TestBroadcastClient(serverAddrFlag, serverPortFlag)
		// client.TestRosteringClient(serverAddrFlag, serverPortFlag)
		// client.TestIncidentReportClient(serverAddrFlag, serverPortFlag)
		// client.TestCameraIotClientUser(serverAddrFlag, serverPortFlag)
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

func initSentry() {
	envFilePath := ".env"
	err := godotenv.Load(envFilePath)

	if err != nil {
		fmt.Println(err)
	}

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              os.Getenv("SENTRY_DNS"),
		TracesSampleRate: 1.0,
		AttachStacktrace: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
