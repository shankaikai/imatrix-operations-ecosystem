import { Dispatch } from "react";
import { Broadcast } from "../proto/operations_ecosys_pb";

export default function getFilteredBroadcasts(
  broadcasts: Broadcast[],
  setFilteredBroadcasts: Dispatch<Broadcast[]>,
  search: string,
  selectValue: string,
  filterValue: string
) {
  console.log("getFilteredBroadcasts called");
  // TODO: Add filtering logic
  console.log(broadcasts);
  const filtered = broadcasts.filter((broadcast) =>
    broadcast.getContent().includes(search)
  );
  console.log(filtered);
  setFilteredBroadcasts(filtered);
}
