import { Stack } from "@mantine/core";
import React from "react";
import { Broadcast } from "../../proto/operations_ecosys_pb";
import BroadcastCard from "./BroadcastCard/BroadcastCard";

const mockdata = [
  {
    id: 1,
    imgUrl:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQKXW9nkZ9NMKDopl5m6XfT2jIAJamJWko3VpstBIlFDKy4VTVYOx4TVeZ8SpzD1ZGdlAs&usqp=CAU",
    name: "Test Name",
    message: "This is a test message!!!",
    unread: 0,
  },
  {
    id: 2,
    imgUrl:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQKXW9nkZ9NMKDopl5m6XfT2jIAJamJWko3VpstBIlFDKy4VTVYOx4TVeZ8SpzD1ZGdlAs&usqp=CAU",
    name: "Test Name",
    message: "This is a longgggggggggggggggg test message!!!",
    unread: 2,
  },
  {
    id: 3,
    imgUrl:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQKXW9nkZ9NMKDopl5m6XfT2jIAJamJWko3VpstBIlFDKy4VTVYOx4TVeZ8SpzD1ZGdlAs&usqp=CAU",
    name: "Test Name",
  },
];

interface BroadcastListProps {
  broadcastList: Broadcast[];
}

export default function BroadcastList({ broadcastList }: BroadcastListProps) {
  // TODO: Iterate through broadcast list and populate this
  return (
    <Stack>
      <BroadcastCard />
      <BroadcastCard />
      <BroadcastCard />
      <BroadcastCard />
    </Stack>
  );
}
