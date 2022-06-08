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
  EmployeeEvaluationResponse,
  Filter,
  Roster,
  RosterFilter,
  RosterQuery,
  RosterResponse,
  User,
} from "../proto/operations_ecosys_pb";
import getRosterDates from "./getRosterDates";
import _ from "lodash";
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

interface RosteringGuardsList {
  [key: string]: User.AsObject[][];
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
    console.log("getAIFSColumns called");

    const client = getRosterClient();
    const filter = new RosterFilter();
    filter.setField(RosterFilter.Field.START_TIME);
    const filterDate = new Filter();
    filterDate.setValue(dayjs(selectedDate).format("YYYY-DD-MM 18:00:00"));
    filterDate.setComparison(Filter.Comparisons.EQUAL);
    filter.setComparisons(filterDate);
    const query = new RosterQuery();
    query.addFilters(filter);

    const stream = client.findRosters(query);
    stream.on("data", (response: RosterResponse) => {
      // console.log(response.toObject());
      const responseRoster = response.getRoster()?.toObject();
      const responseAssignedGuard = response
        .getRoster()
        ?.getGuardAssignedList()[0]
        .getGuardAssigned()
        ?.getEmployee()
        ?.toObject();

      setRosterBaskets((prevBaskets) => {
        return [...prevBaskets, responseRoster!];
      });

      setGuardsAssigned((prevGuards) => {
        let newGuards = prevGuards;
        responseRoster &&
          responseAssignedGuard &&
          newGuards[selectedDate.toString()].splice(responseRoster.aifsId, 0, [
            responseAssignedGuard,
          ]);
        return newGuards;
      });
    });
  };

  const getAvailableGuards = () => {
    console.log("getAvailableGuards called");

    const client = getRosterClient();

    const query = new AvailabilityQuery();
    query.setStartTime(dayjs(selectedDate).format("YYYY-DD-MM 18:00:00"));

    const stream = client.getAvailableUsers(query);

    stream.on("data", (response: EmployeeEvaluationResponse) => {
      // console.log(response);
      const employeeResponse = response
        .getEmployee()
        ?.getEmployee()
        ?.toObject();
      setGuardsAssigned((prevGuards) => {
        let newGuards = _.cloneDeep(prevGuards);
        employeeResponse &&
          newGuards[selectedDate.toString()][0].push(employeeResponse);
        return newGuards;
      });
    });
  };

  useEffect(() => {
    updateRosterDates();
  }, [offset]);

  useEffect(() => {
    resetStates(selectedDate.toString());
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

export function submitNewRoster() {}
