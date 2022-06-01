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
import AcknowledgementToolTip from "./AcknowledgementToolTip";

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
          <AcknowledgementToolTip />
          <AcknowledgementToolTip />
          <AcknowledgementToolTip />
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
