import dayjs from "dayjs";
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
  User,
} from "../proto/operations_ecosys_pb";
import { ENVOY_ADDRESS } from "../utils/constant";
import {
  showApproveReportSuccessNotification,
  showUpdateReportSuccessNotification,
} from "./notifications";

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
  createNewReport?: (
    values: UpdateReport,
    setReports: Dispatch<IncidentReport.AsObject[]>,
    userId: number
  ) => Promise<void>;
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
  const [selectedReport, setSelectedReport] =
    useState<IncidentReport.AsObject>();

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
        createNewReport,
      }}
    >
      {children}
    </ReportingContext.Provider>
  );
}

export function getReportingClient(): IncidentReportServicesClient {
  return new IncidentReportServicesClient(ENVOY_ADDRESS, null);
}

export function useReporting() {
  return useContext(ReportingContext);
}

export const updateReports = (
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
    // console.log(response.getIncidentReport()?.toObject());
    //@ts-ignore
    setReports((oldState) => [
      ...oldState,
      response.getIncidentReport()?.toObject(),
    ]);
  });
};

export interface UpdateReport {
  title: string;
  address: string;
  time: string;
  description: string;
  isPoliceNotified: boolean;
  hasStolenItem: boolean;
}

export async function submitUpdateReport(
  values: UpdateReport,
  id: number,
  setReports: Dispatch<IncidentReport.AsObject[]>,
  userId: number
) {
  const client = getReportingClient();

  const incidentReport = new IncidentReport();
  const incidentReportContent = new IncidentReportContent();
  const user = new User();

  user.setUserId(userId); //TODO: swap with actual user's id when logged in
  incidentReportContent.setTitle(values.title);
  incidentReportContent.setAddress(values.address);
  incidentReportContent.setIncidentTime(
    dayjs(values.time, "YYYY-MM-DD[T]HH:mm").format("YYYY-MM-DD HH:mm")
  );
  incidentReportContent.setDescription(values.description);
  incidentReportContent.setIsPoliceNotified(values.isPoliceNotified);
  incidentReportContent.setHasStolenItem(values.hasStolenItem);

  incidentReport.setLastModifedUser(user);
  incidentReport.setIncidentReportId(id);
  incidentReport.setIncidentReportContent(incidentReportContent);

  await client
    .updateIncidentReport(incidentReport, {})
    .then((response) => {
      console.log(response);
      showUpdateReportSuccessNotification();
    })
    .catch((e) => console.log(e));

  updateReports(0, setReports, true);
}

export async function approveReport(
  id: number,
  setReports: Dispatch<IncidentReport.AsObject[]>
) {
  const client = getReportingClient();

  const incidentReport = new IncidentReport();
  incidentReport.setIsApproved(true);
  incidentReport.setIncidentReportId(id);
  incidentReport.setApprovalDate(dayjs(Date.now()).format("YYYY-MM-DD HH:mm"));

  await client.updateIncidentReport(incidentReport, {}).then((response) => {
    showApproveReportSuccessNotification();
  });

  updateReports(0, setReports, true);

  return id;
}

export async function createNewReport(
  values: UpdateReport,
  setReports: Dispatch<IncidentReport.AsObject[]>,
  userId: number
) {
  const client = getReportingClient();

  const incidentReport = new IncidentReport();
  const incidentReportContent = new IncidentReportContent();
  const user = new User();

  user.setUserId(userId); //TODO: swap with actual user's id when logged in

  incidentReportContent.setTitle(values.title);
  incidentReportContent.setAddress(values.address);
  incidentReportContent.setIncidentTime(
    dayjs(values.time, "YYYY-MM-DD[T]HH:mm").format("YYYY-MM-DD HH:mm")
  );
  incidentReportContent.setDescription(values.description);
  incidentReportContent.setIsPoliceNotified(values.isPoliceNotified);
  incidentReportContent.setHasStolenItem(values.hasStolenItem);
  incidentReportContent.setInjuryDescription("");
  incidentReportContent.setActionTaken("");
  incidentReportContent.setStolenItemDescription("");
  incidentReportContent.setReportImageLink("");

  incidentReport.setCreator(user);
  incidentReport.setCreationDate(dayjs(Date.now()).format("YYYY-MM-DD HH:mm"));
  // incidentReport.setLastModifiedDate(
  //   dayjs(Date.now()).format("YYYY-MM-DD HH:mm")
  // );
  incidentReport.setIsApproved(false);
  incidentReport.setIsOriginal(true);
  incidentReport.setIncidentReportContent(incidentReportContent);

  await client
    .addIncidentReport(incidentReport, {})
    .then((response) => {
      showUpdateReportSuccessNotification();
    })
    .catch((e) => console.log(e));

  updateReports(0, setReports, true);
}
