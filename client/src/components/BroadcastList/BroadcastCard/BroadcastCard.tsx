import {
  Badge,
  Card,
  Group,
  Text,
  Tooltip,
  useMantineTheme,
} from "@mantine/core";
import React from "react";
import { Broadcast } from "../../../proto/operations_ecosys_pb";

interface BroadcastCardProps {
  broadcast: Broadcast;
}

const mockdata = [
  {
    content: "This is a sample message",
  },
];

export default function BroadcastCard({ broadcast }: BroadcastCardProps) {
  const theme = useMantineTheme();
  const recipients = broadcast.getRecipientsList();

  return (
    <Card
      p="lg"
      shadow="sm"
      sx={{ position: "relative", borderLeft: "5px solid red" }}
    >
      <Group position="apart">
        <Text>{broadcast.getContent()}</Text>
        <Group>
          <Tooltip label={"Xavier, Emily, Hannah"}>
            <Badge variant="filled" color="gray">
              AIFS 1
            </Badge>
          </Tooltip>
          <Tooltip label={"Xavier, Emily, Hannah"}>
            <Badge variant="filled" color="gray">
              AIFS 2
            </Badge>
          </Tooltip>
          <Tooltip label={"Xavier, Emily, Hannah"}>
            <Badge variant="filled" color="green">
              AIFS 3
            </Badge>
          </Tooltip>
        </Group>
      </Group>
      <Text
        size="xs"
        color="dimmed"
        sx={{
          position: "absolute",
          right: theme.spacing.lg,
        }}
      >
        {broadcast.getCreationDate()}
      </Text>
    </Card>
  );
}
