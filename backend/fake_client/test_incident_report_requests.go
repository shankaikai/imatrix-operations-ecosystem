package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"capstone.operations_ecosystem/backend/common"
	pb "capstone.operations_ecosystem/backend/proto"
)

func TestIncidentReportClient(serverAddr *string, serverPort *int) {
	report := CreateFakeIncidentReport(1)
	// pk := InsertIncidentReport(serverAddr, serverPort, report)
	report.IncidentReportId = 1 //pk

	// ConsolidatedFindIncidentReportTest(serverAddr, serverPort)
	ConsolidatedUpdateIncidentReportTest(serverAddr, serverPort, report)
	// DeleteIncidentReport(serverAddr, serverPort, &pb.IncidentReport{IncidentReportId: 5})
}

func InsertIncidentReport(serverAddr *string, serverPort *int, report *pb.IncidentReport) int64 {
	fmt.Println("Inserting IncidentReport:", report.IncidentReportId)
	client, conn := createIncidentReportClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.AddIncidentReport(context.Background(), report)
	if err != nil {
		fmt.Println("InsertIncidentReport ERROR:", err)
		return -1
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}

	return res.PrimaryKey
}

func ConsolidatedFindIncidentReportTest(serverAddr *string, serverPort *int) {
	FindIncidentReportsNoFilter(serverAddr, serverPort)
	FindIncidentReportsIdFilter(serverAddr, serverPort)
	FindIncidentReportsContentIdFilter(serverAddr, serverPort)
	// FindIncidentReportsTypeFilter(serverAddr, serverPort)
	// FindIncidentReportsModifierFilter(serverAddr, serverPort)
	// FindIncidentReportsLastModifiedDateFilter(serverAddr, serverPort)
	// FindIncidentReportsIsApprovedFilter(serverAddr, serverPort)
	// FindIncidentReportsSignatureFilter(serverAddr, serverPort)
	// FindIncidentReportsApprovalDateFilter(serverAddr, serverPort)
	// FindIncidentReportsMultipleFilters(serverAddr, serverPort)
}

func FindIncidentReportsNoFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding reports without filter")
	FindIncidentReportsTest(serverAddr, serverPort, &pb.IncidentReportQuery{Limit: 4, Skip: 0})
}

func FindIncidentReportsIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding reports id 2 filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "2"}
	filter := &pb.IncidentReportFilter{Comparisons: com, Field: pb.IncidentReportFilter_REPORT_ID}

	query := &pb.IncidentReportQuery{Limit: 4, Filters: []*pb.IncidentReportFilter{filter}}

	FindIncidentReportsTest(serverAddr, serverPort, query)
}

func FindIncidentReportsContentIdFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding reports content id 2 filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "2"}
	filter := &pb.IncidentReportFilter{Comparisons: com, Field: pb.IncidentReportFilter_REPORT_CONTENT_ID}

	query := &pb.IncidentReportQuery{Limit: 4, Filters: []*pb.IncidentReportFilter{filter}}

	FindIncidentReportsTest(serverAddr, serverPort, query)
}

func FindIncidentReportsTypeFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding reports type filter")

	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: pb.IncidentReport_INTRUDER.String()}
	filter := &pb.IncidentReportFilter{Comparisons: com, Field: pb.IncidentReportFilter_REPORT_TYPE}

	query := &pb.IncidentReportQuery{Limit: 4, Filters: []*pb.IncidentReportFilter{filter}}

	FindIncidentReportsTest(serverAddr, serverPort, query)
}

func FindIncidentReportsModifierFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding reports modifier filter")

	com := &pb.Filter{Comparison: pb.Filter_CONTAINS, Value: "2"}
	filter := &pb.IncidentReportFilter{Comparisons: com, Field: pb.IncidentReportFilter_MODIFIER}

	query := &pb.IncidentReportQuery{Limit: 4, Filters: []*pb.IncidentReportFilter{filter}}

	FindIncidentReportsTest(serverAddr, serverPort, query)
}

func FindIncidentReportsLastModifiedDateFilter(serverAddr *string, serverPort *int) {
	creationDateMax := time.Now().AddDate(0, 0, -1)

	fmt.Println("Finding reports creation date filter, max date:", creationDateMax)
	com := &pb.Filter{Comparison: pb.Filter_LESSER_EQ, Value: creationDateMax.Format(common.DATETIME_FORMAT)}
	filter := &pb.IncidentReportFilter{Comparisons: com, Field: pb.IncidentReportFilter_LAST_MODIFIED_DATE}

	query := &pb.IncidentReportQuery{Limit: 4, Filters: []*pb.IncidentReportFilter{filter}}

	FindIncidentReportsTest(serverAddr, serverPort, query)
}

func FindIncidentReportsIsApprovedFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding reports approved")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	filter := &pb.IncidentReportFilter{Comparisons: com, Field: pb.IncidentReportFilter_IS_APPROVED}

	query := &pb.IncidentReportQuery{Limit: 4, Filters: []*pb.IncidentReportFilter{filter}}

	FindIncidentReportsTest(serverAddr, serverPort, query)
}

func FindIncidentReportsSignatureFilter(serverAddr *string, serverPort *int) {
	fmt.Println("Finding reports creator id filter")
	com := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "3"}
	filter := &pb.IncidentReportFilter{Comparisons: com, Field: pb.IncidentReportFilter_SIGNATURE}

	query := &pb.IncidentReportQuery{Limit: 4, Filters: []*pb.IncidentReportFilter{filter}}

	FindIncidentReportsTest(serverAddr, serverPort, query)
}

func FindIncidentReportsApprovalDateFilter(serverAddr *string, serverPort *int) {
	creationDateMax := time.Now().AddDate(0, 0, -1)

	fmt.Println("Finding reports approval date filter, max date:", creationDateMax)
	com := &pb.Filter{Comparison: pb.Filter_LESSER_EQ, Value: creationDateMax.Format(common.DATETIME_FORMAT)}
	filter := &pb.IncidentReportFilter{Comparisons: com, Field: pb.IncidentReportFilter_APPROVAL_DATE}

	query := &pb.IncidentReportQuery{Limit: 4, Filters: []*pb.IncidentReportFilter{filter}}

	FindIncidentReportsTest(serverAddr, serverPort, query)
}

func FindIncidentReportsMultipleFilters(serverAddr *string, serverPort *int) {
	fmt.Println("Finding reports mutiple filter")

	reportFilters := make([]*pb.IncidentReportFilter, 0)

	// type
	typeCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: pb.IncidentReport_INTRUDER.String()}
	typeFilter := &pb.IncidentReportFilter{Comparisons: typeCom, Field: pb.IncidentReportFilter_REPORT_TYPE}
	reportFilters = append(reportFilters, typeFilter)

	// creation date
	approvalCom := &pb.Filter{Comparison: pb.Filter_EQUAL, Value: "1"}
	approvalfilter := &pb.IncidentReportFilter{Comparisons: approvalCom, Field: pb.IncidentReportFilter_IS_APPROVED}
	reportFilters = append(reportFilters, approvalfilter)

	query := &pb.IncidentReportQuery{Limit: 4, Filters: reportFilters}
	FindIncidentReportsTest(serverAddr, serverPort, query)
}

func FindIncidentReportsTest(serverAddr *string, serverPort *int, query *pb.IncidentReportQuery) {
	fmt.Println("Finding reports")
	client, conn := createIncidentReportClient(serverAddr, serverPort)
	defer conn.Close()

	stream, err := client.FindIncidentReports(context.Background(), query)
	if err != nil {
		fmt.Println("FindIncidentReportsTest ERROR:", err)
		return
	}

	count := 0
	for {
		reportRes, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("FindIncidentReportsTest ERROR:", err)
		}

		fmt.Println("Client received response:", reportRes.Response.Type)
		if reportRes.Response.Type == pb.Response_ERROR {
			continue
		}

		fmt.Println(count, ":", reportRes.IncidentReport)
		count++
	}
}

func ConsolidatedUpdateIncidentReportTest(serverAddr *string, serverPort *int, report *pb.IncidentReport) {
	// UpdateIncidentReportBasicFields(serverAddr, serverPort, report)
	UpdateIncidentReportContent(serverAddr, serverPort, report)
}

func UpdateIncidentReportBasicFields(serverAddr *string, serverPort *int, report *pb.IncidentReport) {
	updateIncidentReport := pb.IncidentReport{
		IncidentReportId: report.IncidentReportId,
		Type:             pb.IncidentReport_OTHERS,
		IsApproved:       true,
		ApprovalDate:     time.Now().Format(common.DATETIME_FORMAT),
	}

	UpdateIncidentReportTest(serverAddr, serverPort, &updateIncidentReport)
}

func UpdateIncidentReportContent(serverAddr *string, serverPort *int, report *pb.IncidentReport) {
	report.IncidentReportContent.ActionTaken = "updated action taken"
	updateIncidentReport := pb.IncidentReport{
		IncidentReportId:      report.IncidentReportId,
		IncidentReportContent: report.IncidentReportContent,
	}

	UpdateIncidentReportTest(serverAddr, serverPort, &updateIncidentReport)
}

func UpdateIncidentReportTest(serverAddr *string, serverPort *int, report *pb.IncidentReport) {
	fmt.Println("Updating IncidentReport...")
	client, conn := createIncidentReportClient(serverAddr, serverPort)
	defer conn.Close()
	res, err := client.UpdateIncidentReport(context.Background(), report)
	if err != nil {
		fmt.Println("UpdateIncidentReportTest ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}

func DeleteIncidentReport(serverAddr *string, serverPort *int, report *pb.IncidentReport) {
	fmt.Println("Deleting IncidentReport:", report.IncidentReportId)
	client, conn := createIncidentReportClient(serverAddr, serverPort)
	defer conn.Close()

	res, err := client.DeleteIncidentReport(context.Background(), report)
	if err != nil {
		fmt.Println("DeleteIncidentReport ERROR:", err)
		return
	}

	fmt.Println("Client received response:", res.Type)

	if res.Type == pb.Response_ERROR {
		fmt.Println("Client received error response:", res.ErrorMessage)
	}
}

func createIncidentReportClient(serverAddr *string, serverPort *int) (pb.IncidentReportServicesClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *serverAddr, *serverPort), opts...)
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewIncidentReportServicesClient(conn)

	return client, conn
}
