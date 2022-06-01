import { Button, Group, Select, Text, TextInput } from "@mantine/core";
import React, { Dispatch } from "react";
import { Plus, Search } from "tabler-icons-react";
import { useBroadcastClient } from "../../../helpers/useBroadcastClient";

interface BroadcastFilterProps {
  setModalOpen: Dispatch<boolean>;
}

export default function BroadcastFilter({
  setModalOpen,
}: BroadcastFilterProps) {
  const {
    search,
    setSearch,
    selectValue,
    setSelectValue,
    filterValue,
    setFilterValue,
  } = useBroadcastClient();

  const handleNewClick = () => {
    setModalOpen(true);
  };

  return (
    <Group
      spacing="xs"
      position="apart"
      sx={{
        display: "sticky",
      }}
    >
      {setSearch && (
        <TextInput
          icon={<Search size={16} />}
          placeholder="Search"
          radius="lg"
          size="xs"
          style={{ flex: "1 1 auto" }}
          value={search}
          onChange={(event) => setSearch(event.currentTarget.value)}
        />
      )}

      <Button
        size="xs"
        radius="xl"
        leftIcon={<Plus size={14} />}
        onClick={handleNewClick}
      >
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
