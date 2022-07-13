import { Button, Group, ScrollArea, Stack } from "@mantine/core";
import React, { useEffect, useState } from "react";
import getFilteredReports from "../../../helpers/getFilteredReports";
import {
  updateReports,
  useReporting,
} from "../../../helpers/useReportingClient";
import { IncidentReport } from "../../../proto/operations_ecosys_pb";
import ReportSmallCard from "../ReportSmallCard";

export default function ReportsList() {
  const { reports, selectedReport, search, selectValue, setReports } =
    useReporting();

  const [filteredReports, setFilteredReports] = useState<
    IncidentReport.AsObject[]
  >([]);

  const handleLoadMore = () => {
    setReports && updateReports(reports.length, setReports);
  };

  useEffect(() => {
    getFilteredReports(reports, setFilteredReports, search, selectValue);
  }, [reports, selectValue, search]);

  return (
    <ScrollArea
      sx={{
        width: "39%",
        height: "calc(100vh - 132px)",
      }}
    >
      <Stack>
        {filteredReports.map(
          (report) =>
            report.incidentReportContent?.title
              .toLowerCase()
              .includes(search.toLowerCase()) && (
              <ReportSmallCard
                key={report.incidentReportId}
                selected={
                  report.incidentReportId === selectedReport?.incidentReportId
                }
                title={report.incidentReportContent?.title || ""}
                sender={report.creator?.name || ""}
                senderImg={report.creator?.userSecurityImg || ""}
                aifsId={1} //TODO: do we need this?
                id={report.incidentReportId}
                creationDate={report.creationDate}
                updateDate={report.lastModifiedDate}
                report={report}
              />
            )
        )}
      </Stack>
      <Group position="center" mt="lg">
        <Button onClick={handleLoadMore}>Load More</Button>
      </Group>
    </ScrollArea>
  );
}
