import { Roster } from "../proto/operations_ecosys_pb";


// Checks the overall status of the roster 
export default function getOverallRosterStatus(
  rosters: Roster.AsObject[]
): Roster.Status {
  for (var roster of rosters) {
    if (roster.status === Roster.Status.REJECTED) {
      return Roster.Status.REJECTED;
    }
  }

  for (var roster of rosters) {
    if (roster.status === Roster.Status.IS_DEFAULT) {
      return Roster.Status.IS_DEFAULT;
    }
  }

  for (var roster of rosters) {
    if (roster.status === Roster.Status.PENDING) {
      return Roster.Status.PENDING;
    }
  }

  return Roster.Status.CONFIRMED;
}
