import React from "react";

interface AcknowledgementToolTipProps {
  AIFS: BroadcastRecipient[];
}

import { useState } from "react";
import {
  Popover,
  Badge,
  Image,
  Text,
  Stack,
  Group,
  Avatar,
  Indicator,
} from "@mantine/core";
import { BroadcastRecipient } from "../../../../proto/operations_ecosys_pb";

export default function AcknowledgementToolTip() {
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
          color="gray"
        >
          AIFS 1
        </Badge>
      }
    >
      <Stack>
        <Group position="apart">
          <Text size="sm">AIFS 1</Text>
          <Text size="sm">Location</Text>
        </Group>
        <Group>
          <Stack align="center" spacing="xs">
            <Indicator color="green" offset={5} withBorder size={14}>
              <Avatar radius="xl" />
            </Indicator>
            <Text size="xs">Name</Text>
            <Text size="xs">HP: 91119999</Text>
          </Stack>
          <Stack align="center" spacing="xs">
            <Indicator color="green" offset={5} withBorder size={14}>
              <Avatar radius="xl" />
            </Indicator>
            <Text size="xs">Name</Text>
            <Text size="xs">HP: 91119999</Text>
          </Stack>
          <Stack align="center" spacing="xs">
            <Indicator color="green" offset={5} withBorder size={14}>
              <Avatar
                radius="xl"
                src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=250&q=80"
              />
            </Indicator>
            <Text size="xs">Name</Text>
            <Text size="xs">HP: 91119999</Text>
          </Stack>
        </Group>
      </Stack>
    </Popover>
  );
}
