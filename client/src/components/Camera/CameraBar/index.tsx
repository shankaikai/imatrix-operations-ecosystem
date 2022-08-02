import { Group, TextInput } from "@mantine/core";
import React from "react";
import { Search } from "tabler-icons-react";

export default function CameraBar() {
  return (
    <Group>
      <TextInput
        data-testid="cameraSearch"
        icon={<Search size={16} />}
        placeholder="Search"
        radius="lg"
        size="xs"
        style={{ flex: "1 1 auto" }}
      />
    </Group>
  );
}
