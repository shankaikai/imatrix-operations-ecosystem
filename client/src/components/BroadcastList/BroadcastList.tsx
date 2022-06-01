import {
  Button,
  Group,
  ScrollArea,
  Stack,
  useMantineTheme,
} from "@mantine/core";
import React, { useEffect, useState } from "react";
import getFilteredBroadcasts from "../../helpers/getFilteredBroadcasts";
import { useBroadcastClient } from "../../helpers/useBroadcastClient";
import { Broadcast } from "../../proto/operations_ecosys_pb";
import BroadcastCard from "./BroadcastCard/BroadcastCard";

interface BroadcastListProps {}

export default function BroadcastList({}: BroadcastListProps) {
  const { broadcasts, search, selectValue, filterValue, setBroadcasts } =
    useBroadcastClient();

  const handleLoadMoreClick = () => {
    // console.log(broadcasts);
    console.log(filteredBroadcasts);
  };

  const [filteredBroadcasts, setFilteredBroadcasts] = useState<Broadcast[]>([]);

  // Run the filtering function whenever the filter options changes
  useEffect(() => {
    getFilteredBroadcasts(
      broadcasts,
      setFilteredBroadcasts,
      search,
      selectValue,
      filterValue
    );
  }, [broadcasts, search, selectValue, filterValue]);

  // mockBroadcast.setCreationDate(Date.now());

  const b = new Broadcast();
  b.setContent("This is mock broadcast");
  const c = new Broadcast();
  c.setContent("This is mock broadcast 2");

  // TODO: Iterate through broadcast list and filter
  return (
    <ScrollArea
      sx={{
        height: "calc(100vh - 60px - 30px - 30px)",
      }}
    >
      <Stack>
        {filteredBroadcasts.map((broadcast) => (
          <BroadcastCard
            key={broadcast.getBroadcastId()}
            broadcast={broadcast}
          />
        ))}
        {/* <BroadcastCard broadcast={b} />
        <BroadcastCard broadcast={b} />
        <BroadcastCard broadcast={b} />
        <BroadcastCard broadcast={b} />
        <BroadcastCard broadcast={b} />
        <BroadcastCard broadcast={b} />
        <BroadcastCard broadcast={b} />
        <BroadcastCard broadcast={b} />
        <BroadcastCard broadcast={b} /> */}

        <Group position="center">
          <Button onClick={handleLoadMoreClick}>Load More</Button>
        </Group>
      </Stack>
    </ScrollArea>
  );
}
