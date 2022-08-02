import { Modal, ScrollArea, Stack } from "@mantine/core";
import type { NextPage } from "next";
import { useState } from "react";
import BroadcastFilter from "../../components/Broadcasting/BroadcastFilter/BroadcastFilter";
import BroadcastList from "../../components/Broadcasting/BroadcastList/BroadcastList";
import NewBroadcast from "../../components/Broadcasting/NewBroadcast";
import Layout from "../../components/Layout/Layout";
import { BroadcastProvider } from "../../helpers/useBroadcastClient";

const Broadcasting: NextPage = () => {
  // TODO: Convert into context provider
  const [modalOpen, setModalOpen] = useState(false);

  return (
    <Layout>
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
    </Layout>
  );
};

export default Broadcasting;
