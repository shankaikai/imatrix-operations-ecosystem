import dayjs from "dayjs";
import { Dispatch } from "react";
import { RosteringGuardsList } from "./useRosteringClient";

// When an available guard is dragged onto a 
export default function addGuardToGuardsAssigned(
  id: number,
  index: number,
  aifsId: number,
  selectedDate?: Date,
  setGuardsAssigned?: Dispatch<RosteringGuardsList>
) {
  selectedDate &&
    setGuardsAssigned &&
    setGuardsAssigned((prevState) => {
      let newState = _.cloneDeep(prevState);
      const date = dayjs(selectedDate).format("YYYY-MM-DD");
      const [orginalGuard] = newState[date][aifsId].splice(0, 1);
      const [newGuard] = newState[date][0].splice(index, 1, orginalGuard);
      newState[date][aifsId].splice(0, 0, newGuard);
      return newState;
    });
}
