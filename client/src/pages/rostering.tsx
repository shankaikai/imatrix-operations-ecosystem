import { Group, ScrollArea, Stack, Text } from "@mantine/core";
import type { NextPage } from "next";
import { useState } from "react";
import {
  DragDropContext,
  Draggable,
  DropResult,
  Droppable,
} from "react-beautiful-dnd";
import RosterAvailability from "../components/Rostering/RosterAvailability";
import RosterBasket from "../components/Rostering/RosterBasket";
import RosterDateBar from "../components/Rostering/RosterDateBar";
import RosterGuard from "../components/Rostering/RosterGuard";
import { RosteringProvider } from "../helpers/useRosteringClient";

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
  const [aifs1, setAifs1] = useState([]);

  const onDragEnd = (result: DropResult) => {
    const { source, destination } = result;
    if (!destination) return;

    const items = Array.from(aifs1);
    const [newOrder] = items.splice(source.index, 1);
    items.splice(destination.index, 0, newOrder);

    setAifs1(items);
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
            <ScrollArea sx={{ width: "40%" }}>
              <Stack>
                {aifs.map((aif) => (
                  <RosterBasket key={aif} title={aif} />
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
