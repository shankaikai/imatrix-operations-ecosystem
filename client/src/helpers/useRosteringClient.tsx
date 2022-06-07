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
  Filter,
  Roster,
  RosterFilter,
  RosterQuery,
  RosterResponse,
} from "../proto/operations_ecosys_pb";
import getRosterDates from "./getRosterDates";

interface RosteringContextInterface {
  rosterDates: Date[];
  setRosterDates?: Dispatch<Date[]>;
  offset: number;
  setOffset?: Dispatch<number>;
  selectedDate?: Date;
  setSelectedDate?: Dispatch<Date>;
  rosterBaskets?: Roster.AsObject[];
}

const RosteringContext = createContext<RosteringContextInterface>({
  rosterDates: [],
  offset: 0,
});

interface RosteringProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function RosteringProvider({ children }: RosteringProviderProps) {
  const [rosterDates, setRosterDates] = useState<Date[]>([]);
  const [offset, setOffset] = useState<number>(0);
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());

  const [rosterBaskets, setRosterBaskets] = useState<Roster.AsObject[]>([]);

  const updateRosterDates = () => {
    const dates = getRosterDates(offset);
    console.log("dates:", dates);
    setRosterDates(dates);
  };

  // Get basket data
  const getAIFSColumns = () => {
    console.log("getAIFSColumns called");
    const client = getRosterClient();
    const filter = new RosterFilter();
    filter.setField(RosterFilter.Field.START_TIME);
    const filterDate = new Filter();
    filterDate.setValue(dayjs(selectedDate).format("YYYY-DD-MM 18:00:00"));
    filter.setComparisons(filterDate);
    const query = new RosterQuery();
    query.addFilters(filter);
    const stream = client.findRosters(query);
    stream.on("data", (response: RosterResponse) => {
      setRosterBaskets((oldState) => [
        ...oldState,
        response.toObject().roster!,
      ]);
    });
  };

  useEffect(() => {
    updateRosterDates();
  }, [offset]);

  useEffect(() => {
    getAIFSColumns();
  }, [selectedDate]);

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
