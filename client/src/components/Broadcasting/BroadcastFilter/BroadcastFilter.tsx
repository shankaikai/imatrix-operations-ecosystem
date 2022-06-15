import { ActionIcon, Button, Group, Select, Text, TextInput } from "@mantine/core";
import React, { Dispatch } from "react";
import { Plus, Search, Refresh } from "tabler-icons-react";
import { useBroadcast } from "../../../helpers/useBroadcastClient";

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
    updateBroadcasts,
    setBroadcasts
  } = useBroadcast();

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
          data-testid="broadcastSearch"
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
        data-testid="newBroadcastButton"
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
          data-testid="sortBy"
          data={[
            { value: "latest", label: "Latest First" },
            { value: "oldest", label: "Oldest First" },
          ]}
          value={selectValue}
          //@ts-ignore
          onChange={setSelectValue}
          variant="unstyled"
          size="xs"
          sx={{
            width: "100px",
          }}
        />
      </Group>
      <Group>
        <Text size="xs" color="dimmed">
          Filter by:
        </Text>
        <Select
          data-testid="filterBy"
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
          sx={{
            width: "100px",
          }}
        />
      </Group>
      <Group>
        <ActionIcon onClick={()=>{(updateBroadcasts && setBroadcasts) && updateBroadcasts(0, setBroadcasts , true)}}>
          <Refresh size={14} />
        </ActionIcon>
      </Group>
    </Group>
  );
}
