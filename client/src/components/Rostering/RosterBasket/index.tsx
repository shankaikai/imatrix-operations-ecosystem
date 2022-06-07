import { Card, Group, Stack, Text, useMantineTheme } from "@mantine/core";
import React, { useState } from "react";
import { Droppable } from "react-beautiful-dnd";
import { Roster } from "../../../proto/operations_ecosys_pb";

interface RosterCardProps {
  basket: Roster.AsObject;
  index: number;
}

export default function RosterBasket({ basket, index }: RosterCardProps) {
  const theme = useMantineTheme();

  return (
    <Card>
      <Stack>
        <Group>
          <Text>{`AIFS ${basket.aifsId}`}</Text>
        </Group>
        <Droppable droppableId={index.toString()}>
          {(provided, snapshot) => (
            <div
              style={{
                border: `2px dashed ${
                  theme.colorScheme === "dark"
                    ? theme.colors.dark[3]
                    : theme.colors.gray[5]
                }`,
                borderStyle: "dashed",
                height: "100px",
                borderRadius: theme.spacing.sm,
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
              }}
              {...provided.droppableProps}
              ref={provided.innerRef}
            >
              <Text color="dimmed">Drag and drop the I-Specialists here</Text>
            </div>
          )}
        </Droppable>
      </Stack>
    </Card>
  );
}
