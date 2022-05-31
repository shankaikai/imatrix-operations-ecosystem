import { Button, Group, ScrollArea, Stack } from "@mantine/core";
import React from "react";
import { useBroadcastClient } from "../../helpers/useBroadcastClient";
import { Broadcast } from "../../proto/operations_ecosys_pb";
import BroadcastCard from "./BroadcastCard/BroadcastCard";

interface BroadcastListProps {}

export default function BroadcastList({}: BroadcastListProps) {
  const { broadcasts } = useBroadcastClient();

  const handleLoadMoreClick = () => {
    console.log(broadcasts);
  };

  // TODO: Iterate through broadcast list and populate this
  return (
    <ScrollArea>
      <Stack>
        {broadcasts.map((broadcast) => (
          <BroadcastCard
            key={broadcast.getBroadcastId()}
            broadcast={broadcast}
          />
        ))}
        {/* <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} />
        <BroadcastCard broadcast={new Broadcast()} /> */}
        <Group position="center">
          <Button onClick={handleLoadMoreClick}>Load More</Button>
        </Group>
      </Stack>
    </ScrollArea>
  );
}
