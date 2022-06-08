import { Dispatch } from "react";
import { RosteringGuardsList } from "./useRosteringClient";

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
      const date = selectedDate.toString();
      const [orginalGuard] = newState[date][aifsId].splice(0, 1);
      const [newGuard] = newState[date][0].splice(index, 1, orginalGuard);
      newState[date][aifsId].splice(0, 0, newGuard);
      return newState;
    });
}
