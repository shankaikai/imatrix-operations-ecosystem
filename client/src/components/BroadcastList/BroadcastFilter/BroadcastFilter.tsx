import {
  Group,
  TextInput,
  Select,
  createStyles,
  Text,
  Button,
} from "@mantine/core";
import { Plus, Search } from "tabler-icons-react";
import React, { Dispatch, useState } from "react";

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

interface BroadcastFilterProps {
  search: string;
  setSearch: Dispatch<string>;
  selectValue: string;
  setSelectValue: Dispatch<string>;
  filterValue: string;
  setFilterValue: Dispatch<string>;
}

export default function BroadcastFilter({
  search,
  setSearch,
  selectValue,
  setSelectValue,
  filterValue,
  setFilterValue,
}: BroadcastFilterProps) {
  const { classes } = useStyles();
  return (
    <Group className={classes.filters} position="apart">
      <TextInput
        icon={<Search size={16} />}
        placeholder="Search"
        radius="lg"
        size="xs"
        style={{ flex: "1 1 auto" }}
        value={search}
        onChange={(event) => setSearch(event.currentTarget.value)}
      />
      <Button size="xs" radius="xl" leftIcon={<Plus size={14} />}>
        New
      </Button>
      <Group>
        <Text size="xs" color="dimmed">
          Sort by:
        </Text>
        <Select
          data={[
            { value: "latest", label: "Lastest First" },
            { value: "oldest", label: "Oldest First" },
          ]}
          value={selectValue}
          //@ts-ignore
          onChange={setSelectValue}
          variant="unstyled"
          size="xs"
        />
      </Group>
      <Group>
        <Text size="xs" color="dimmed">
          Filter by:
        </Text>
        <Select
          data={[
            { value: "all", label: "All" },
            { value: "1", label: "AIFS 1" },
            { value: "2", label: "AIFS 2" },
            { value: "3", label: "AIFS 3" },
          ]}
          value={filterValue}
          //@ts-ignore
          onChange={setFilterValue}
          variant="unstyled"
          size="xs"
        />
      </Group>
    </Group>
  );
}
