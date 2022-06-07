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
  rosterBaskets: Roster[];
  rosterDates: Date[];
  setRosterDates?: Dispatch<Date[]>;
  offset: number;
  setOffset?: Dispatch<number>;
  selectedDate?: Date;
  setSelectedDate?: Dispatch<Date>;
}

const RosteringContext = createContext<RosteringContextInterface>({
  rosterDates: [],
  offset: 0,
  rosterBaskets: [],
});

interface RosteringProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function RosteringProvider({ children }: RosteringProviderProps) {
  const [rosterDates, setRosterDates] = useState<Date[]>([]);
  const [offset, setOffset] = useState<number>(0);
  const [selectedDate, setSelectedDate] = useState<Date>(new Date());

  const [rosterBaskets, setRosterBaskets] = useState<Roster[]>([]);
  const [availableBaskets, setAvailableBaskets] = useState();

  const updateRosterDates = () => {
    const dates = getRosterDates(offset);
    setRosterDates(dates);
  };

  // Get basket data
  const updateRosterBaskets = () => {
    console.log("getAIFSColumns called");

    setRosterBaskets(() => []);

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
      console.log(response.toObject());
      const responseRoster = response.getRoster();

      setRosterBaskets((prevBaskets) => {
        console.log(prevBaskets);
        return [...prevBaskets, responseRoster!];
      });
    });
  };

  useEffect(() => {
    updateRosterDates();
  }, [offset]);

  useEffect(() => {
    updateRosterBaskets();
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
