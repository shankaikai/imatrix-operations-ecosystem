import { Card, Group, Text, useMantineTheme } from "@mantine/core";
import React from "react";
import { AIFSRecipient } from "../../../helpers/recipientsFormatter";
import AcknowledgementToolTip from "./AcknowledgementToolTip";

interface BroadcastCardProps {
  content: string;
  date: string;
  aifs: AIFSRecipient[];
}

export default function BroadcastCard({
  content,
  date,
  aifs,
}: BroadcastCardProps) {
  const theme = useMantineTheme();
  let allAcknowledged = true;
  for (var recipient of aifs) {
    allAcknowledged = allAcknowledged && recipient.allAcknowledged;
  }

  return (
    <Card
      p="lg"
      shadow="sm"
      sx={{
        position: "relative",
        borderLeft: `5px solid ${allAcknowledged ? "green" : "red"}`,
      }}
    >
      <Group position="apart">
        <Text>{content}</Text>
        <Group>
          {aifs.map((aif) => (
            <AcknowledgementToolTip key={aif.id} data={aif} />
          ))}
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
        {date}
      </Text>
    </Card>
  );
}
