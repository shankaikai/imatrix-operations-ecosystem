import { Button, Group, Input, Text, ActionIcon, Select } from "@mantine/core";
import React from "react";
import { GridDots, GridPattern, List, Plus, Search } from "tabler-icons-react";

export default function ReportsBar() {
  return (
    <Group position="apart">
      <Text>Reports</Text>
      <Input
        radius="xl"
        variant="filled"
        placeholder="Search incident reports"
        icon={<Search size={14} />}
      />
      <Button radius="xl" leftIcon={<Plus size={14} />}>
        New
      </Button>
      <Group>
        <ActionIcon variant="light">
          <List />
        </ActionIcon>
        <ActionIcon variant="light">
          <GridDots />
        </ActionIcon>
        <Text color="dimmed" size="sm">
          Sort by:
        </Text>
        <Select variant="filled" data={["Latest First", "Oldest First"]} />
      </Group>
    </Group>
  );
}
