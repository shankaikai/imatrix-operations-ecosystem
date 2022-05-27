import { Group } from "@mantine/core";
import type { NextPage } from "next";
import { useEffect, useState } from "react";
import BroadcastList from "../components/BroadcastList";
import BroadcastMessages from "../components/BroadcastMessages";
import useBroadcastClient from "../helpers/useBroadcastClient";
import {
  BroadcastQuery,
  BroadcastResponse,
  Broadcast,
} from "./../proto/operations_ecosys_pb";

const Home: NextPage = () => {
  const [broadcast, setBroadcast] = useState<Broadcast>();
  const [selectedCard, setSelectedCard] = useState<number>(-1);

  const broadcastClient = useBroadcastClient();

  const getBroadcasts = async () => {
    const query = new BroadcastQuery();
    var stream = broadcastClient.findBroadcasts(query, {});

    stream.on("data", (response: BroadcastResponse) => {
      console.log(response.getBroadcast());
      setBroadcast(response.getBroadcast());
    });
  };

  useEffect(() => {
    getBroadcasts();
  }, []);

  return (
    <Group spacing={0}>
      {/* <p>{broadcast?.getContent()}</p> */}
      <BroadcastList
        selectedCard={selectedCard}
        setSelectedCard={setSelectedCard}
      />
      <BroadcastMessages selectedCard={selectedCard} />
    </Group>
  );
};

export default Home;
