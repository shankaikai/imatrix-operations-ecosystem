import { Group, Stack, Text, useMantineTheme } from "@mantine/core";
import React from "react";
import { useRostering } from "../../../helpers/useRosteringClient";
import RosterGuard from "../RosterGuard";

export default function RosterAvailability() {
  const theme = useMantineTheme();
  const { guardsAssigned, selectedDate } = useRostering();

  console.log("guardsAssigned from RosterAvailability", guardsAssigned);

  return (
    <Stack>
      <Text>Available</Text>
      <Group>
        {selectedDate &&
          guardsAssigned[selectedDate.toString()] &&
          guardsAssigned[selectedDate.toString()][0].map((guard, index) => (
            <RosterGuard key={guard.userId} guard={guard} index={index} />
          ))}
      </Group>
    </Stack>
  );
}
