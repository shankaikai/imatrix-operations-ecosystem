import {
  Button,
  Group,
  ScrollArea,
  Stack,
  useMantineTheme,
} from "@mantine/core";
import dayjs from "dayjs";
import React, { useEffect, useState } from "react";
import getFilteredBroadcasts from "../../../helpers/getFilteredBroadcasts";
import recipientFormatter from "../../../helpers/recipientsFormatter";
import { useBroadcast } from "../../../helpers/useBroadcastClient";
import { Broadcast } from "../../../proto/operations_ecosys_pb";
import BroadcastCard from "../BroadcastCard/BroadcastCard";

interface BroadcastListProps {}

export default function BroadcastList({}: BroadcastListProps) {
  const { broadcasts, search, selectValue, filterValue, updateBroadcasts } =
    useBroadcast();

  const handleLoadMoreClick = () => {
    // console.log(broadcasts);
    // console.log(filteredBroadcasts);
    updateBroadcasts && updateBroadcasts();
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

  return (
    <ScrollArea
      sx={{
        height: "calc(100vh - 60px - 30px - 30px)",
      }}
    >
      <Stack data-testid="broadcastList">
        {filteredBroadcasts.map((broadcast) => {
          const key = broadcast.getBroadcastId();
          const content = broadcast.getContent();
          const date = dayjs(
            broadcast.getCreationDate()?.toDate() as Date
          ).format("DD/MM/YYYY, h:mm A");
          const aifs = recipientFormatter(broadcast.getRecipientsList());
          return (
            <BroadcastCard
              key={key}
              content={content}
              date={date}
              aifs={aifs}
            />
          );
        })}

        <Group position="center">
          <Button onClick={handleLoadMoreClick}>Load More</Button>
        </Group>
      </Stack>
    </ScrollArea>
  );
}
