import {
  createContext,
  Dispatch,
  useContext,
  useEffect,
  useState,
} from "react";
import { IncidentReportServicesClient } from "../proto/Operations_ecosysServiceClientPb";
import {
  IncidentReport,
  IncidentReportContent,
  IncidentReportQuery,
  IncidentReportResponse,
} from "../proto/operations_ecosys_pb";

interface RosteringContextInterface {
  reports: IncidentReport.AsObject[];
  setReports?: Dispatch<IncidentReport.AsObject[]>;
  search: string;
  setSearch?: Dispatch<string>;
  selectValue: string;
  setSelectValue?: Dispatch<string>;
  modalOpen: boolean;
  setModalOpen?: Dispatch<boolean>;
  selectedReport?: IncidentReport.AsObject;
  setSelectedReport?: Dispatch<IncidentReport.AsObject>;
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
});

interface ReportingProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function ReportingProvider({ children }: ReportingProviderProps) {
  const [reports, setReports] = useState<IncidentReport.AsObject[]>([]);
  const [search, setSearch] = useState("");
  const [selectValue, setSelectValue] = useState("latest");
  const [modalOpen, setModalOpen] = useState(false);
  const [selectedReport, setSelectedReport] = useState<IncidentReport.AsObject>();

  useEffect(() => {
    updateReports(reports.length, setReports);
  }, []);

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

const updateReports = (
  skip: number,
  setReports: Dispatch<IncidentReport.AsObject[]>,
  clear?: boolean
) => {
  console.log("updateReports called");

  //@ts-ignore
  clear && setReports(() => []);

  const client = getReportingClient();

  const query = new IncidentReportQuery();
  query.setSkip(skip);

  const stream = client.findIncidentReports(query);

  // On every data received, add it to the state
  stream.on("data", (response: IncidentReportResponse) => {
    console.log(response.getIncidentReport()?.toObject());
    //@ts-ignore
    setReports((oldState) => [
      ...oldState,
      response.getIncidentReport()?.toObject(),
    ]);
  });
};
