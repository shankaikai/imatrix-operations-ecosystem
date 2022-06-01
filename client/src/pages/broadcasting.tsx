import { Modal, ScrollArea, Stack } from "@mantine/core";
import type { NextPage } from "next";
import { useState } from "react";
import BroadcastFilter from "../components/BroadcastList/BroadcastFilter/BroadcastFilter";
import BroadcastList from "../components/BroadcastList/BroadcastList";
import NewBroadcast from "../components/BroadcastList/NewBroadcast";
import { BroadcastProvider } from "../helpers/useBroadcastClient";

const Broadcasting: NextPage = () => {
  // TODO: Convert into context provider
  const [search, setSearch] = useState<string>("");
  const [selectValue, setSelectValue] = useState("latest");
  const [filterValue, setFilterValue] = useState("all");
  const [modalOpen, setModalOpen] = useState(false);

  return (
    <BroadcastProvider>
      <Modal
        opened={modalOpen}
        onClose={() => setModalOpen(false)}
        title="Broadcast Details"
      >
        <NewBroadcast setModelOpen={setModalOpen} />
      </Modal>
      <Stack spacing="xs">
        <BroadcastFilter setModalOpen={setModalOpen} />
        <BroadcastList />
      </Stack>
    </BroadcastProvider>
  );
};

export default Broadcasting;
