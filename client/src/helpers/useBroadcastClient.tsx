import {
  createContext,
  Dispatch,
  useContext,
  useEffect,
  useState,
} from "react";
import { BroadcastServicesClient } from "../proto/Operations_ecosysServiceClientPb";
import {
  AIFSBroadcastRecipient,
  Broadcast,
  BroadcastQuery,
  BroadcastResponse,
  User,
} from "../proto/operations_ecosys_pb";
import {
  showErrorNotification,
  showBroadcastSuccessNotification,
} from "./notifications";

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
  const [search, setSearch] = useState("");
  const [selectValue, setSelectValue] = useState("latest");
  const [filterValue, setFilterValue] = useState("all");

  const updateBroadcasts = () => {
    console.log("updateBroadcasts called");
    const client = getBroadcastClient();

    const query = new BroadcastQuery();
    query.setSkip(broadcasts.length);

    const stream = client.findBroadcasts(query);

    // On every data received, add it to the state
    stream.on("data", (response: BroadcastResponse) => {
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

export function useBroadcast() {
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

  var recipientList: AIFSBroadcastRecipient[] = [];

  console.log(recipient);

  if (recipient.includes("all")) {
    for (var id of [1, 2, 3]) {
      const broadcastRecipient = new AIFSBroadcastRecipient();
      broadcastRecipient.setAifsId(id);
      recipientList.push(broadcastRecipient);
    }
  } else {
    for (var user of recipient) {
      const broadcastRecipient = new AIFSBroadcastRecipient();
      broadcastRecipient.setAifsId(parseInt(user));
      recipientList.push(broadcastRecipient);
    }
  }

  console.log(recipientList);
  // TODO: Discuss recipient logic
  broadcast.setRecipientsList(recipientList);
  broadcast.setUrgency(urgencyMap[urgency[0]]);
  broadcast.setContent(message);
  broadcast.setType(Broadcast.BroadcastType.ANNOUNCEMENT);
  // TODO: Change to user context
  const creator = new User();
  creator.setUserId(2);
  broadcast.setCreator(creator);

  console.log(broadcast);

  await client
    .addBroadcast(broadcast, {})
    .then((response) => {
      showBroadcastSuccessNotification();
      console.log(response);
    })
    .catch((error) => {
      // TODO: Error toast
      showErrorNotification();
    });
}
