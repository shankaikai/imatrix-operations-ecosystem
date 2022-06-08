import { Group, ScrollArea, Stack, Text } from "@mantine/core";
import React from "react";
import { useRostering } from "../../../helpers/useRosteringClient";
import RosterAvailability from "../RosterAvailability";
import RosterBasketsLists from "../RosterBasketsList";
import RosterDateBar from "../RosterDateBar";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";

export default function RosterContainer() {
  const { guardsAssigned, setGuardsAssigned } = useRostering();

  return (
    <DndProvider backend={HTML5Backend}>
      <Stack>
        <RosterDateBar />
        <Text size="xl" weight={500}>
          Rostering
        </Text>
        <Group position="apart" align="flex-start">
          <ScrollArea sx={{ width: "45%" }}>
            <RosterBasketsLists />
          </ScrollArea>
          <ScrollArea sx={{ width: "40%" }}>
            <RosterAvailability />
          </ScrollArea>
        </Group>
      </Stack>
    </DndProvider>
  );
}
