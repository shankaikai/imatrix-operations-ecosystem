import dayjs from "dayjs";
import {
  createContext,
  Dispatch,
  useContext,
  useEffect,
  useState,
} from "react";
import { RosterServicesClient } from "../proto/Operations_ecosysServiceClientPb";
import {
  AvailabilityQuery,
  BulkRosters,
  EmployeeEvaluation,
  EmployeeEvaluationResponse,
  Filter,
  Roster,
  RosterAssignement,
  RosterFilter,
  RosterQuery,
  RosterResponse,
  User,
} from "../proto/operations_ecosys_pb";
import getRosterDates from "./getRosterDates";
import _ from "lodash";
import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb";

interface RosteringContextInterface {
  rosterBaskets: Roster.AsObject[];
  rosterDates: Date[];
  setRosterDates?: Dispatch<Date[]>;
  offset: number;
  setOffset?: Dispatch<number>;
  selectedDate?: Date;
  setSelectedDate?: Dispatch<Date>;
  guardsAssigned: RosteringGuardsList;
  setGuardsAssigned?: Dispatch<RosteringGuardsList>;
}

const RosteringContext = createContext<RosteringContextInterface>({
  rosterDates: [],
  offset: 0,
  rosterBaskets: [],
  guardsAssigned: {},
});

interface RosteringProviderProps {
  children: JSX.Element | JSX.Element[];
}

export interface RosteringGuardsList {
  [key: string]: EmployeeEvaluation.AsObject[][];
}
export function RosteringProvider({ children }: RosteringProviderProps) {
  const [rosterDates, setRosterDates] = useState<Date[]>([]);
  const [offset, setOffset] = useState<number>(0);
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());

  const [rosterBaskets, setRosterBaskets] = useState<Roster.AsObject[]>([]);

  const [guardsAssigned, setGuardsAssigned] = useState<RosteringGuardsList>({});
  // const [guardsAssigned, setGuardsAssigned] = useState<User.AsObject[][]>([[]]);
  // Date bar
  const updateRosterDates = () => {
    const dates = getRosterDates(offset);
    setRosterDates(dates);
  };

  const resetStates = (date: string) => {
    setRosterBaskets(() => []);

    let guardsAssignedCopy = _.cloneDeep(guardsAssigned);
    guardsAssignedCopy[date] = [[]];
    setGuardsAssigned(guardsAssignedCopy);
  };

  // Get basket data
  const updateRosterBaskets = () => {
    console.log("updateRosterBaskets called");

    const client = getRosterClient();
    const filter = new RosterFilter();
    filter.setField(RosterFilter.Field.START_TIME);
    const filterDate = new Filter();
    filterDate.setValue(dayjs(selectedDate).format("YYYY-MM-DD 18:00:00"));
    filterDate.setComparison(Filter.Comparisons.EQUAL);
    filter.setComparisons(filterDate);
    const query = new RosterQuery();
    query.addFilters(filter);

    const stream = client.findRosters(query);

    stream.on("data", (response: RosterResponse) => {
      console.log(response.toObject());
      const responseRoster = response.getRoster()?.toObject();
      const responseAssignedGuard = response
        .getRoster()
        ?.getGuardAssignedList()[0]
        .getGuardAssigned()
        ?.toObject();

      setRosterBaskets((prevBaskets) => {
        return [...prevBaskets, responseRoster!];
      });

      setGuardsAssigned((prevGuards) => {
        let newGuards = prevGuards;
        console.log(prevGuards);
        responseRoster &&
          responseAssignedGuard &&
          newGuards[dayjs(selectedDate).format("YYYY-MM-DD")].splice(
            responseRoster.aifsId,
            0,
            [responseAssignedGuard]
          );
        return newGuards;
      });
    });
  };

  const getAvailableGuards = () => {
    console.log("getAvailableGuards called");

    const client = getRosterClient();

    const query = new AvailabilityQuery();
    query.setStartTime(dayjs(selectedDate).format("YYYY-MM-DD 18:00:00"));

    const stream = client.getAvailableUsers(query);

    stream.on("data", (response: EmployeeEvaluationResponse) => {
      // console.log(response);
      const employeeResponse = response.getEmployee()?.toObject();
      setGuardsAssigned((prevGuards) => {
        let newGuards = _.cloneDeep(prevGuards);
        employeeResponse &&
          newGuards[dayjs(selectedDate).format("YYYY-MM-DD")][0].push(
            employeeResponse
          );
        return newGuards;
      });
    });
  };

  useEffect(() => {
    updateRosterDates();
  }, [offset]);

  useEffect(() => {
    resetStates(dayjs(selectedDate).format("YYYY-MM-DD"));
    updateRosterBaskets();
    getAvailableGuards();
  }, [selectedDate]);

  // useEffect(() => {
  //   console.log(guardsAssigned);
  // }, [guardsAssigned]);

  return (
    <RosteringContext.Provider
      value={{
        rosterDates,
        setRosterDates,
        offset,
        setOffset,
        selectedDate,
        setSelectedDate,
        rosterBaskets,
        guardsAssigned,
        setGuardsAssigned,
      }}
    >
      {children}
    </RosteringContext.Provider>
  );
}

export function getRosterClient(): RosterServicesClient {
  // TODO: add the envoy address into .env
  return new RosterServicesClient("http://localhost:8080", null, {});
}

export function useRostering() {
  return useContext(RosteringContext);
}

export function submitNewRoster(
  guardsAssigned: RosteringGuardsList,
  date: Date
) {
  const client = getRosterClient();

  const rosterList: Roster[] = [];

  for (const i of [1, 2, 3]) {
    const userObject = guardsAssigned[dayjs(date).format("YYYY-MM-DD")][i][0];

    const user = new User();
    userObject.employee && user.setUserId(userObject.employee.userId);

    const employeeEvaluation = new EmployeeEvaluation();
    employeeEvaluation.setEmployee(user);

    const rosterAssignment = new RosterAssignement();
    rosterAssignment.setGuardAssigned(employeeEvaluation);

    const roster = new Roster();
    roster.addGuardAssigned(rosterAssignment);
    roster.setAifsId(i);

    const timeStampStart = new Timestamp();
    const startTime = new Date();
    startTime.setDate(date.getMonth());
    startTime.setMonth(date.getMonth());
    startTime.setFullYear(date.getFullYear());
    startTime.setHours(26, 0, 0);
    timeStampStart.fromDate(startTime);

    const timeStampEnd = new Timestamp();
    const endTime = new Date();

    roster.setStartTime(dayjs(date).format("YYYY-MM-DD 18:00:00"));
    const endDate = new Date();
    endDate.setDate(date.getDate() + 1);
    roster.setEndTime(dayjs(endDate).format("YYYY-MM-DD 18:00:00"));

    rosterList.push(roster);
  }

  const bulkRoster = new BulkRosters();
  bulkRoster.setRostersList(rosterList);

  client
    .addRoster(bulkRoster, {})
    .then((response) => {
      // TODO: Show success toast
      console.log(response);
    })
    .catch((err) => {
      console.log(err);
    });
}
