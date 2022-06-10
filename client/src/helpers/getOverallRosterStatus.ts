import { Roster } from "../proto/operations_ecosys_pb";

export default function getOverallRosterStatus(
  rosters: Roster.AsObject[]
): Roster.Status {
  for (var roster of rosters) {
    if (roster.status === Roster.Status.REJECTED) {
      return Roster.Status.REJECTED;
    }
  }

  for (var roster of rosters) {
    if (roster.status === Roster.Status.PENDING) {
      return Roster.Status.PENDING;
    }
  }

  for (var roster of rosters) {
    if (roster.isDefault) {
      return Roster.Status.PENDING;
    }
  }


  return Roster.Status.CONFIRMED;
}
