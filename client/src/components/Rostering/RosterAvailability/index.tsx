import {
  Stack,
  Text,
  useMantineTheme
} from "@mantine/core";
import React from "react";
import { Draggable, Droppable } from "react-beautiful-dnd";
import { useRostering } from "../../../helpers/useRosteringClient";
import RosterGuard from "../RosterGuard";

export default function RosterAvailability() {
  const theme = useMantineTheme();
  const { guardsAssigned } = useRostering();

  return (
    <Stack>
      <Text>Available</Text>
      <Droppable droppableId="0">
        {(provided, snapshot) => (
          <div
            {...provided.droppableProps}
            ref={provided.innerRef}
            style={{
              display: "flex",
              flexDirection: "row",
              flexWrap: "wrap",
            }}
          >
            {guardsAssigned[0].map((guard, index) => (
              <Draggable
                key={guard.userId}
                draggableId={guard.userId.toString()}
                index={index}
              >
                {(provided, snapshot) => (
                  <div
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}
                    ref={provided.innerRef}
                    style={{
                      ...provided.draggableProps.style,
                      padding: theme.spacing.xs,
                    }}
                  >
                    <RosterGuard guard={guard} />
                  </div>
                )}
              </Draggable>
            ))}
            {provided.placeholder}
          </div>
        )}
      </Droppable>
    </Stack>
  );
}
