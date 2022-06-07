import {
  createContext,
  Dispatch,
  useContext,
  useEffect,
  useState,
} from "react";
import getRosterDates from "./getRosterDates";

interface RosteringContextInterface {
  rosterDates: Date[];
  setRosterDates?: Dispatch<Date[]>;
  offset: number;
  setOffset?: Dispatch<number>;
  selectedDate?: number;
  setSelectedDate?: Dispatch<number>;
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
  const [selectedDate, setSelectedDate] = useState<number>(
    new Date().getDate()
  );

  const [droppableDivs, setDroppableDivs] = useState([
    
  ]);

  const updateRosterDates = () => {
    const dates = getRosterDates(offset);
    setRosterDates(dates);
  };

  useEffect(() => {
    updateRosterDates();
  }, [offset]);

  return (
    <RosteringContext.Provider
      value={{
        rosterDates,
        setRosterDates,
        offset,
        setOffset,
        selectedDate,
        setSelectedDate,
      }}
    >
      {children}
    </RosteringContext.Provider>
  );
}

export function useRostering() {
  return useContext(RosteringContext);
}
