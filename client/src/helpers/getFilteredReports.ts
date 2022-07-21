import { Dispatch } from "react";
import { IncidentReport } from "../proto/operations_ecosys_pb";

export default function getFilteredReports(
  reports: IncidentReport.AsObject[],
  setFilteredReports: Dispatch<IncidentReport.AsObject[]>,
  search: string,
  selectValue: string
) {
  console.log("getFilteredReports called");
  // TODO: Add filtering logic
  let filtered = reports.filter((report) => {
    const title = report.incidentReportContent?.title as string;
    const includesSearch = title.toLowerCase().includes(search.toLowerCase());
    return includesSearch;
  });

  if (selectValue == "oldest") {
    filtered = filtered.reverse();
  }

  setFilteredReports(filtered);
}
