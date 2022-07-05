import { createContext, Dispatch, useContext, useState } from "react";
import { IncidentReportServicesClient } from "../proto/Operations_ecosysServiceClientPb";
import { IncidentReport } from "../proto/operations_ecosys_pb";

interface RosteringContextInterface {
  reports: IncidentReport.AsObject[];
  setReports?: Dispatch<IncidentReport.AsObject[]>;
  search: string;
  setSearch?: Dispatch<string>;
  selectValue: string;
  setSelectValue?: Dispatch<string>;
  modalOpen: boolean;
  setModalOpen?: Dispatch<boolean>;
  selectedReport: number;
  setSelectedReport?: Dispatch<number>;
  updateReports?: (
    skip: number,
    setRosters: Dispatch<IncidentReport.AsObject[]>,
    clear?: boolean
  ) => void;
}

const ReportingContext = createContext<RosteringContextInterface>({
  reports: [],
  search: "",
  selectValue: "latest",
  modalOpen: false,
  selectedReport: -1,
});

interface ReportingProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function ReportingProvider({ children }: ReportingProviderProps) {
  const [reports, setReports] = useState<IncidentReport.AsObject[]>([]);
  const [search, setSearch] = useState("");
  const [selectValue, setSelectValue] = useState("latest");
  const [modalOpen, setModalOpen] = useState(false);
  const [selectedReport, setSelectedReport] = useState(-1);

  return (
    <ReportingContext.Provider
      value={{
        reports,
        setReports,
        search,
        setSearch,
        selectValue,
        setSelectValue,
        modalOpen,
        setModalOpen,
        selectedReport,
        setSelectedReport,
      }}
    >
      {children}
    </ReportingContext.Provider>
  );
}

export function getReportingClient(): IncidentReportServicesClient {
  return new IncidentReportServicesClient("http://localhost:8080", null);
}

export function useReporting() {
  return useContext(ReportingContext);
}
