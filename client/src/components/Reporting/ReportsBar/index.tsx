import {
  Button,
  Group,
  Text,
  ActionIcon,
  Select,
  TextInput,
} from "@mantine/core";
import React from "react";
import { GridDots, GridPattern, List, Plus, Search } from "tabler-icons-react";
import { useReporting } from "../../../helpers/useReportingClient";

export default function ReportsBar() {
  const { search, setSearch, selectValue, setSelectValue, setModalOpen } =
    useReporting();

  const handleNewClick = () => {
    console.log("handleNewClick called");
    setModalOpen && setModalOpen(true);
  };

  return (
    <Group position="apart">
      <Text>Reports</Text>
      {setSearch && (
        <TextInput
          radius="xl"
          placeholder="Search incident reports"
          icon={<Search size={14} />}
          sx={{
            flex: "auto 1",
          }}
          value={search}
          onChange={(event) => setSearch(event.currentTarget.value)}
        />
      )}
      <Button
        radius="xl"
        leftIcon={<Plus size={14} />}
        onClick={handleNewClick}
      >
        New
      </Button>
      <Group>
        <Text color="dimmed" size="sm">
          Sort by:
        </Text>
        <Select
          variant="filled"
          data={[
            { label: "Latest First", value: "latest" },
            { label: "Oldest First", value: "oldest" },
          ]}
          value={selectValue}
          onChange={setSelectValue}
        />
      </Group>
    </Group>
  );
}
