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
  updateBroadcasts?: (skip: number, setBroadcasts: Dispatch<Broadcast[]>, clear?:boolean) => void;
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

const updateBroadcasts = (skip: number, setBroadcasts: Dispatch<Broadcast[]>, clear?: boolean) => {
  console.log("updateBroadcasts called");

  //@ts-ignore
  clear && setBroadcasts(() => []);

  const client = getBroadcastClient();

  const query = new BroadcastQuery();
  query.setSkip(skip);

  const stream = client.findBroadcasts(query);

  // On every data received, add it to the state
  stream.on("data", (response: BroadcastResponse) => {
    //@ts-ignore
    setBroadcasts((oldState) => [...oldState, response.getBroadcast()!]);
  });
};

export function BroadcastProvider({ children }: BroadcastProviderProps) {
  const [broadcasts, setBroadcasts] = useState<Broadcast[]>([]);
  const [search, setSearch] = useState("");
  const [selectValue, setSelectValue] = useState("latest");
  const [filterValue, setFilterValue] = useState("all");

  

  // Call once when first render
  useEffect(() => {
    updateBroadcasts(broadcasts.length, setBroadcasts);
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
  setBroadcasts
}: {
  recipient: string[];
  urgency: string;
  message: string;
  setBroadcasts: Dispatch<Broadcast[]>
}) {
  const client = getBroadcastClient();

  const broadcast = new Broadcast();

  interface urgencyMapInterface {
    [urgency: string]: Broadcast.UrgencyType;
  }

  // TODO: Ask gab wats good practice for this
  const urgencyMap: urgencyMapInterface = {
    "Not Urgent": Broadcast.UrgencyType.LOW,
    Urgent: Broadcast.UrgencyType.HIGH,
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
  broadcast.setUrgency(urgencyMap[urgency]);
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
      updateBroadcasts(0, setBroadcasts, true)
    })
    .catch((error) => {
      // TODO: Error toast
      showErrorNotification();
    });
}
