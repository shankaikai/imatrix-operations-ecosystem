import { Card, Group, Stack, Text, useMantineTheme } from "@mantine/core";
import React from "react";
import { Draggable, Droppable } from "react-beautiful-dnd";
import { useRostering } from "../../../helpers/useRosteringClient";
import { Roster } from "../../../proto/operations_ecosys_pb";
import RosterGuard from "../RosterGuard";

interface RosterCardProps {
  basket: Roster.AsObject;
  index: number;
}

export default function RosterBasket({ basket, index }: RosterCardProps) {
  const theme = useMantineTheme();
  const { guardsAssigned } = useRostering();
  // console.log("guardsAssigned", guardsAssigned);

  return (
    <Card>
      <Stack>
        <Group>
          <Text>{`AIFS ${basket.aifsId}`}</Text>
        </Group>
        <Droppable droppableId={basket.aifsId.toString()}>
          {(provided, snapshot) => (
            <div
              style={{
                border: `2px dashed ${
                  theme.colorScheme === "dark"
                    ? theme.colors.dark[3]
                    : theme.colors.gray[5]
                }`,
                borderStyle: "dashed",
                height: "auto",
                padding: theme.spacing.xs,
                borderRadius: theme.spacing.sm,
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
              }}
              {...provided.droppableProps}
              ref={provided.innerRef}
            >
              {guardsAssigned[basket.aifsId] ? (
                <Draggable
                  key={guardsAssigned[basket.aifsId][0].userId}
                  draggableId={guardsAssigned[
                    basket.aifsId
                  ][0].userId.toString()}
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
                        guard={guardsAssigned[basket.aifsId][0]}
                        withLabels
                      />
                    </div>
                  )}
                </Draggable>
              ) : (
                <Text>hi</Text>
              )}
              {provided.placeholder}
            </div>
          )}
        </Droppable>
      </Stack>
    </Card>
  );
}
