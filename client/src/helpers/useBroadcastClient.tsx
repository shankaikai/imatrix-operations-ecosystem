import {
  createContext,
  Dispatch,
  useContext,
  useEffect,
  useState,
} from "react";
import { BroadcastServicesClient } from "../proto/Operations_ecosysServiceClientPb";
import {
  Broadcast,
  BroadcastQuery,
  BroadcastResponse,
} from "../proto/operations_ecosys_pb";

interface BroadcastContextInterface {
  broadcasts: Broadcast[];
  setBroadcasts?: Dispatch<Broadcast[]>;
  updateBroadcasts?: () => void;
}

const BroadcastContext = createContext<BroadcastContextInterface>({
  broadcasts: [],
  setBroadcasts: undefined,
  updateBroadcasts: undefined,
});

interface BroadcastProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function BroadcastProvider({ children }: BroadcastProviderProps) {
  const [broadcasts, setBroadcasts] = useState<Broadcast[]>([]);

  const updateBroadcasts = () => {
    console.log("updateBroadcasts called");
    const client = getBroadcastClient();

    const query = new BroadcastQuery();
    query.setSkip(broadcasts.length);

    const stream = client.findBroadcasts(query, {});

    // On every data received, add it to the state
    stream.on("data", (response) => {
      // console.log(response.getBroadcast());
      setBroadcasts((oldState) => [...oldState, response.getBroadcast()!]);
    });
  };

  // Call once when first render
  useEffect(() => {
    updateBroadcasts();
  }, []);

  return (
    <BroadcastContext.Provider
      value={{ broadcasts, setBroadcasts, updateBroadcasts }}
    >
      {children}
    </BroadcastContext.Provider>
  );
}

export function getBroadcastClient(): BroadcastServicesClient {
  // TODO: add the envoy address into .env
  return new BroadcastServicesClient("http://localhost:8080", null, {});
}

export function useBroadcastClient() {
  return useContext(BroadcastContext);
}
