import { Group, Stack, Text, useMantineTheme } from "@mantine/core";
import React from "react";
import { useRostering } from "../../../helpers/useRosteringClient";
import RosterGuard from "../RosterGuard";

export default function RosterAvailability() {
  const { guardsAssigned, selectedDate } = useRostering();

  // console.log("guardsAssigned from RosterAvailability", guardsAssigned);

  return (
    <Stack>
      <Text>Available</Text>
      <Group>
        {selectedDate &&
          guardsAssigned[selectedDate.toString()] &&
          guardsAssigned[selectedDate.toString()][0].map((guard, index) => (
            <RosterGuard
              key={guard.employee?.userId}
              guard={guard.employee}
              index={index}
            />
          ))}
      </Group>
    </Stack>
  );
}
