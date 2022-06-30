import { Avatar, Card, Group, Stack, Text } from "@mantine/core";
import React from "react";

export default function ReportSmallCard() {
  return (
    <Card shadow="lg" p="xs">
      <Stack>
        <Text weight={500}>Camera Faulty.doc</Text>
        <Group spacing="xs">
          <Avatar radius="xl" />
          <Text size="xs" color="dimmed">
            Philip Wee (AIFS1)
          </Text>
          <Text>•</Text>
          <Text size="xs" color="dimmed">
            Updated on 7/6/22 at 19:04
          </Text>
        </Group>
      </Stack>
    </Card>
  );
}
