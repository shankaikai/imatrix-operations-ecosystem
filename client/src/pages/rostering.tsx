import { Group, ScrollArea, Stack, Text } from "@mantine/core";
import type { NextPage } from "next";
import { useEffect } from "react";
import { DragDropContext, DropResult } from "react-beautiful-dnd";
import RosterAvailability from "../components/Rostering/RosterAvailability";
import RosterBasketsLists from "../components/Rostering/RosterBasketsLists";
import RosterDateBar from "../components/Rostering/RosterDateBar";
import { RosteringProvider, useRostering } from "../helpers/useRosteringClient";

// const aifs = ["AIFS 1 (AMKC)", "AIFS 2 (BKP)", "AIFS 3 (PKC)"];

// const guards = [
//   {
//     id: 1,
//     name: "Guard1",
//     img: "https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg",
//     phone: "92818838",
//   },
//   {
//     id: 2,
//     name: "Guard2",
//     img: "https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg",
//     phone: "92818838",
//   },
//   {
//     id: 3,
//     name: "Guard3",
//     img: "https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg",
//     phone: "92818838",
//   },
// ];

const Rostering: NextPage = () => {
  const onDragEnd = (result: DropResult) => {
    const { source, destination } = result;
    if (!destination) return;
  };

  return (
    <RosteringProvider>
      <DragDropContext onDragEnd={(result) => console.log(result)}>
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
      </DragDropContext>
    </RosteringProvider>
  );
};

export default Rostering;
