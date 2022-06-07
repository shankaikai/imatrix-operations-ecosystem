import { Group, ScrollArea, Stack, Text } from "@mantine/core";
import type { NextPage } from "next";
import { useEffect } from "react";
import {
  DragDropContext,
  Draggable,
  Droppable,
  DropResult,
} from "react-beautiful-dnd";
import RosterBasket from "../components/Rostering/RosterBasket";
import RosterDateBar from "../components/Rostering/RosterDateBar";
import RosterGuard from "../components/Rostering/RosterGuard";
import { RosteringProvider, useRostering } from "../helpers/useRosteringClient";

const aifs = ["AIFS 1 (AMKC)", "AIFS 2 (BKP)", "AIFS 3 (PKC)"];

const guards = [
  {
    id: 1,
    name: "Guard1",
    img: "https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg",
    phone: "92818838",
  },
  {
    id: 2,
    name: "Guard2",
    img: "https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg",
    phone: "92818838",
  },
  {
    id: 3,
    name: "Guard3",
    img: "https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg",
    phone: "92818838",
  },
];

const Rostering: NextPage = () => {
  const { rosterBaskets } = useRostering();

  const onDragEnd = (result: DropResult) => {
    const { source, destination } = result;
    if (!destination) return;
  };

  useEffect(() => {
    console.log("rosterBaskets:", rosterBaskets);
  }, [rosterBaskets]);

  return (
    <RosteringProvider>
      <DragDropContext onDragEnd={(result) => console.log(result)}>
        <Stack>
          <RosterDateBar />
          <Text size="xl" weight={500}>
            Rostering
          </Text>
          <Group position="apart" align="flex-start">
            <ScrollArea sx={{ width: "50%" }}>
              <Stack>
                {rosterBaskets.map((basket, index) => (
                  <RosterBasket
                    key={basket.getAifsId()}
                    basket={basket.toObject()}
                    index={index}
                  />
                ))}
              </Stack>
            </ScrollArea>
            <Droppable droppableId="guardStack">
              {(provided, snapshot) => (
                <div {...provided.droppableProps} ref={provided.innerRef}>
                  {guards.map((guard, index) => (
                    <Draggable
                      key={guard.id.toString()}
                      draggableId={guard.id.toString()}
                      index={index}
                    >
                      {(provided, snapshot) => (
                        <div
                          {...provided.draggableProps}
                          {...provided.dragHandleProps}
                          ref={provided.innerRef}
                          style={{
                            ...provided.draggableProps.style,
                          }}
                        >
                          <RosterGuard
                            id={guard.id}
                            name={guard.name}
                            img={guard.img}
                            phone={guard.phone}
                          />
                        </div>
                      )}
                    </Draggable>
                  ))}
                </div>
              )}
            </Droppable>

            {/* <RosterAvailability /> */}
          </Group>
        </Stack>
      </DragDropContext>
    </RosteringProvider>
  );
};

export default Rostering;
