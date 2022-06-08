import { Group, ScrollArea, Stack, Text } from "@mantine/core";
import type { NextPage } from "next";
import React from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";
import RosterAvailability from "../components/Rostering/RosterAvailability";
import RosterBasketsLists from "../components/Rostering/RosterBasketsList";
import RosterDateBar from "../components/Rostering/RosterDateBar";
import { RosteringProvider } from "../helpers/useRosteringClient";

const Rostering: NextPage = () => {
  return (
    <RosteringProvider>
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
    </RosteringProvider>
  );
};

export default Rostering;
