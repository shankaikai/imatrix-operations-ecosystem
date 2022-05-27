import { Group, TextInput, Select, createStyles, Text } from "@mantine/core";
import { Search } from "tabler-icons-react";
import React, { useState } from "react";

const useStyles = createStyles((theme) => ({
  tabcontainer: {
    height: "100vh",
    width: "360px",
    padding: theme.spacing.xs,
  },
  filters: {
    columnGap: theme.spacing.xs,
  },
  stack: {
    marginTop: theme.spacing.xs,
  },
}));

export default function BroadcastFilter() {
  const [value, setValue] = useState("latest");
  const [search, setSearch] = useState("");

  const { classes } = useStyles();
  return (
    <Group className={classes.filters} position="apart">
      <TextInput
        icon={<Search size={16} />}
        placeholder="Search"
        radius="lg"
        size="xs"
        style={{ width: "170px" }}
        value={search}
        onChange={(event) => setSearch(event.currentTarget.value)}
      />
      <Group>
        <Text size="xs" color="dimmed">
          Sort by:
        </Text>
        <Select
          data={[
            { value: "latest", label: "Lastest First" },
            { value: "oldest", label: "Oldest First" },
          ]}
          value={value}
          //@ts-ignore
          onChange={setValue}
          variant="unstyled"
          size="xs"
          style={{ width: "95px" }}
        />
      </Group>
    </Group>
  );
}
