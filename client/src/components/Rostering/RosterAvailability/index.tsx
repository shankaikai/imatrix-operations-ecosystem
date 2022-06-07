import { Avatar, Group, Select, Stack } from "@mantine/core";
import React, { useState } from "react";

export default function RosterAvailability() {
  const [value, setValue] = useState("Available");

  return (
    <Stack
      sx={{
        width: "55%",
        height: "100%",
      }}
      justify="flex-start"
    >
      <Group>
        <Select
          data={["Available", "Not Available"]}
          value={value}
          //@ts-ignore
          onChange={setValue}
          variant="unstyled"
        />
      </Group>
      <Group>
        <Avatar
          size={80}
          src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
          radius={100}
        />
        <Avatar
          size={80}
          src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
          radius={100}
        />
        <Avatar
          size={80}
          src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
          radius={100}
        />
        <Avatar
          size={80}
          src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
          radius={100}
        />
        <Avatar
          size={80}
          src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
          radius={100}
        />
        <Avatar
          size={80}
          src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
          radius={100}
        />
        <Avatar
          size={80}
          src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
          radius={100}
        />
        <Avatar
          size={80}
          src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
          radius={100}
        />
        <Avatar
          size={80}
          src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
          radius={100}
        />
      </Group>
    </Stack>
  );
}
