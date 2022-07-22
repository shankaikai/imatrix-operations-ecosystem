import {
  Avatar,
  Badge,
  Group,
  Indicator,
  Popover,
  Stack,
  Text,
} from "@mantine/core";
import React, { useState } from "react";
import { AIFSRecipient } from "../../../../helpers/recipientsFormatter";

interface AcknowledgementToolTipProps {
  data: AIFSRecipient;
}

export default function AcknowledgementToolTip({
  data,
}: AcknowledgementToolTipProps) {
  const [opened, setOpened] = useState(false);

  return (
    <Popover
      opened={opened}
      onClose={() => setOpened(false)}
      position="bottom"
      placement="end"
      withArrow
      trapFocus={false}
      closeOnEscape={false}
      transition="pop-top-left"
      styles={{ body: { pointerEvents: "none" } }}
      target={
        <Badge
          onMouseEnter={() => setOpened(true)}
          onMouseLeave={() => setOpened(false)}
          variant="filled"
          color={data.allAcknowledged ? "green" : "gray"}
        >
          {data.id}
        </Badge>
      }
    >
      <Stack>
        <Group position="apart">
          <Text size="sm">{data.id}</Text>
          <Text size="sm">{data.location}</Text>
        </Group>
        <Group>
          {data.users.map((user) => (
            <Stack key={user.id} align="center" spacing="xs">
              <Indicator
                color={user.acknowledged ? "green" : "red"}
                offset={5}
                withBorder
                size={14}
              >
                <Avatar radius="xl" src={user.img} />
              </Indicator>
              <Text size="xs">{user.name}</Text>
              <Text size="xs">HP: {user.phone}</Text>
            </Stack>
          ))}
        </Group>
      </Stack>
    </Popover>
  );
}
