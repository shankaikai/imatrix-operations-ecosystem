import { resolve } from "path";
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
  BroadcastRecipient,
} from "../proto/operations_ecosys_pb";

interface BroadcastContextInterface {
  broadcasts: Broadcast[];
  setBroadcasts?: Dispatch<Broadcast[]>;
  search: string;
  setSearch?: Dispatch<string>;
  selectValue: string;
  setSelectValue?: Dispatch<string>;
  filterValue: string;
  setFilterValue?: Dispatch<string>;
  updateBroadcasts?: () => void;
}

const BroadcastContext = createContext<BroadcastContextInterface>({
  broadcasts: [],
  search: "",
  selectValue: "latest",
  filterValue: "all",
});

interface BroadcastProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function BroadcastProvider({ children }: BroadcastProviderProps) {
  const [broadcasts, setBroadcasts] = useState<Broadcast[]>([]);
  const [search, setSearch] = useState<string>("");
  const [selectValue, setSelectValue] = useState("latest");
  const [filterValue, setFilterValue] = useState("all");

  const updateBroadcasts = () => {
    console.log("updateBroadcasts called");
    const client = getBroadcastClient();

    const query = new BroadcastQuery();
    query.setSkip(broadcasts.length);

    const stream = client.findBroadcasts(query, {});

    // On every data received, add it to the state
    stream.on("data", (response) => {
      console.log(response.getBroadcast());
      setBroadcasts((oldState) => [...oldState, response.getBroadcast()!]);
    });
  };

  // Call once when first render
  useEffect(() => {
    updateBroadcasts();
  }, []);

  return (
    <BroadcastContext.Provider
      value={{
        broadcasts,
        setBroadcasts,
        updateBroadcasts,
        search,
        setSearch,
        selectValue,
        setSelectValue,
        filterValue,
        setFilterValue,
      }}
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

export async function submitNewBroadcast({
  recipient,
  urgency,
  message,
}: {
  recipient: string[];
  urgency: string[];
  message: string;
}) {
  const client = getBroadcastClient();

  const broadcast = new Broadcast();

  interface urgencyMapInterface {
    [urgency: string]: Broadcast.UrgencyType;
  }

  // TODO: Ask gab wats good practice for this
  const urgencyMap: urgencyMapInterface = {
    Low: Broadcast.UrgencyType.LOW,
    Medium: Broadcast.UrgencyType.MEDIUM,
    High: Broadcast.UrgencyType.HIGH,
  };

  var recipientList: BroadcastRecipient[] = [];

  if (recipient.includes("all")) {
    for (var id of [1, 2, 3]) {
      const broadcastRecipient = new BroadcastRecipient();
      broadcastRecipient.setAifsId(id);
      recipientList.push(broadcastRecipient);
    }
  } else {
    for (var user of recipient) {
      const broadcastRecipient = new BroadcastRecipient();
      broadcastRecipient.setAifsId(parseInt(user));
      recipientList.push(broadcastRecipient);
    }
  }

  // TODO: Discuss recipient logic
  broadcast.setRecipientsList(recipientList);
  broadcast.setUrgency(urgencyMap[urgency[0]]);
  broadcast.setContent(message);

  await client
    .addBroadcast(broadcast, {})
    .then((response) => {
      // TODO: Bring up success toast
      console.log(response);
    })
    .catch((error) => {
      // TODO: Error toast
      console.log(error);
    });
}
