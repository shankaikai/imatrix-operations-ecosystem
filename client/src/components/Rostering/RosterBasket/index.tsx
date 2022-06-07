import { Card, Group, Stack, Text, useMantineTheme } from "@mantine/core";
import React, { useState } from "react";
import { Droppable } from "react-beautiful-dnd";

interface RosterCardProps {
  title: string;
}

export default function RosterBasket({ title }: RosterCardProps) {
  const theme = useMantineTheme();
  const [count, setCount] = useState(0);
  const [basket, setBasket] = useState([]);

  return (
    <Card>
      <Stack>
        <Group position="apart">
          <Text>{title}</Text>
          <Text>{count}/3</Text>
        </Group>
        <Droppable droppableId={title}>
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
