import { ScrollArea, Stack } from "@mantine/core";
import React from "react";
import { useReporting } from "../../../helpers/useReportingClient";
import ReportSmallCard from "../ReportSmallCard";

export default function ReportsList() {
  const { reports, selectedReport } = useReporting();
  return (
    <ScrollArea
      sx={{
        width: "39%",
        height: "calc(100vh - 132px)",
      }}
    >
      <Stack>
        {reports.map((report) => (
          <ReportSmallCard
            key={report.incidentReportId}
            selected={
              report.incidentReportId === selectedReport?.incidentReportId
            }
            title={report.incidentReportContent?.title || ""}
            sender={report.creator?.name || ""}
            senderImg={report.creator?.userSecurityImg || ""}
            aifsId={1}
            id={report.incidentReportId}
            creationDate={report.creationDate}
            updateDate={report.lastModifiedDate}
            report={report}
          />
        ))}
      </Stack>
    </ScrollArea>
  );
}
