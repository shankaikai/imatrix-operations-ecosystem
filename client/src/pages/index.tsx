import type { NextPage } from "next";
import { useEffect, useState } from "react";
import useBroadcastClient from "../helpers/useBroadcastClient";
import {
  BroadcastQuery,
  BroadcastResponse,
  Broadcast,
} from "./../proto/operations_ecosys_pb";

const Home: NextPage = () => {
  return (
    <div>
      <h1>Home</h1>
    </div>
  );
};

export default Home;
