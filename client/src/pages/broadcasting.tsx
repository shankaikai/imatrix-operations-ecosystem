import { Stack } from "@mantine/core";
import type { NextPage } from "next";
import { useEffect, useState } from "react";
import BroadcastFilter from "../components/BroadcastList/BroadcastFilter/BroadcastFilter";
import BroadcastList from "../components/BroadcastList/BroadcastList";
import useBroadcastClient from "../helpers/useBroadcastClient";
import {
  Broadcast,
  BroadcastQuery,
  BroadcastResponse,
} from "./../proto/operations_ecosys_pb";

const Broadcasting: NextPage = () => {
  // TODO: Convert into context provider
  const [broadcast, setBroadcast] = useState<Broadcast[]>([]);

  const [search, setSearch] = useState<string>("");
  const [selectValue, setSelectValue] = useState("latest");
  const [filterValue, setFilterValue] = useState("all");

  const broadcastClient = useBroadcastClient();

  // TODO: Shift into helper and store into broadcast context
  const getBroadcasts = async () => {
    const query = new BroadcastQuery();
    console.log("here");
    var stream = broadcastClient.findBroadcasts(query, {});

    stream.on("data", (response: BroadcastResponse) => {
      console.log("ddata");

      console.log(response.getBroadcast());
      setBroadcast([...broadcast, response.getBroadcast() as Broadcast]);
    });
  };

  useEffect(() => {
    getBroadcasts();
  }, []);

  return (
    <Stack spacing="xs">
      <BroadcastFilter
        search={search}
        setSearch={setSearch}
        selectValue={selectValue}
        setSelectValue={setSelectValue}
        filterValue={filterValue}
        setFilterValue={setFilterValue}
      />
      <BroadcastList broadcastList={broadcast} />
    </Stack>
  );
};

export default Broadcasting;
