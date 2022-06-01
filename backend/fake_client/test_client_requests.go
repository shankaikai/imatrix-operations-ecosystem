package client

import (
	"context"
	"fmt"
	"io"

	pb "capstone.operations_ecosystem/backend/proto"
)

func TestAdminClientClient(serverAddr *string, serverPort *int) {
	client := createFakeClient(1)
	pk := InsertClient(serverAddr, serverPort, client)
	client.ClientId = pk
	ConsolidatedFindClientTest(serverAddr, serverPort)
	UpdateClientTest(serverAddr, serverPort, client)
	DeleteClientTest(serverAddr, serverPort, &pb.Client{ClientId: 90})
}

func InsertClient(serverAddr *string, serverPort *int, client *pb.Client) int64 {
	fmt.Println("Inserting client:", client.Name)
	grpcClient, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := grpcClient.AddClient(context.Background(), client)
	if err != nil {
		fmt.Println("InsertClient ERROR:", err)
		return -1
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}

	return res.PrimaryKey
}

func ConsolidatedFindClientTest(serverAddr *string, serverPort *int) {
	FindClientsNoFilter(serverAddr, serverPort)
	FindClientIdFilter(serverAddr, serverPort)
	// FindClientMultipleFilters(serverAddr, serverPort)
}

func FindClientsNoFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding clients without filter")
	FindClientsTest(serverAddr, serverPort, &pb.ClientQuery{Limit: 5})
}

func FindClientIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding client id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.ClientFilter{Comparisons: com, Field: pb.ClientFilter_CLIENT_ID}

	query := &pb.ClientQuery{Limit: 4, Filters: []*pb.ClientFilter{filter}}

	FindClientsTest(serverAddr, serverPort, query)
}

func FindClientsMultipleFilters(serverAddr *string, serverPort *int) {
	fmt.Println("Finding client mutiple filter")

	clientFilters := make([]*pb.ClientFilter, 0)

	// Add filters here when there are more filters

	query := &pb.ClientQuery{Limit: 4, Filters: clientFilters}
	FindClientsTest(serverAddr, serverPort, query)
}

func FindClientsTest(serverAddr *string, serverPort *int, query *pb.ClientQuery) {
	fmt.Println("Finding clients...")
	grpcClient, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	stream, err := grpcClient.FindClients(context.Background(), query)

	if err != nil {
		fmt.Println("FindClientsNoFilter ERROR:", err)
		return
	}

	count := 0
	for {
		clientRes, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("FindClientsNoFilter ERROR:", err)
		}

		fmt.Println("Client received response:", clientRes.Response.Type)
		if clientRes.Response.Type == pb.Response_ERROR {
			continue
		}

		fmt.Println(count, ":", clientRes.Client)
		count++
	}
}

func UpdateClientTest(serverAddr *string, serverPort *int, client *pb.Client) {
	updatedClient := &pb.Client{
		ClientId:             client.ClientId,
		Name:                 "Some new name",
		NumberOfGuardsNeeded: 4,
	}

	fmt.Println("Updating client:", client.ClientId)

	grpcClient, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := grpcClient.UpdateClient(context.Background(), updatedClient)
	if err != nil {
		fmt.Println("UpdateClientTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}

func DeleteClientTest(serverAddr *string, serverPort *int, client *pb.Client) {
	fmt.Println("Deleting client:", client.Name)
	grpcClient, conn := createAdminClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := grpcClient.DeleteClient(context.Background(), client)
	if err != nil {
		fmt.Println("DeleteClientTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}
