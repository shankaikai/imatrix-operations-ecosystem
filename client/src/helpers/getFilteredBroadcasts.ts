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
  let filtered = broadcasts.filter((broadcast) => {
    const content = broadcast.getContent();
    const includesSearch = content.toLowerCase().includes(search.toLowerCase());

    let includesFilter = true;

    if (filterValue !== "all") {
      console.log("not all");
      const recipients = broadcast.getRecipientsList();
      var aifs = [];
      for (var recipient of recipients) {
        aifs.push(recipient.getAifsId().toString());
      }
      includesFilter = aifs.includes(filterValue);
    }
    return includesSearch && includesFilter;
  });

  if (selectValue == "oldest") {
    filtered = filtered.reverse();
  }

  setFilteredBroadcasts(filtered);
}
