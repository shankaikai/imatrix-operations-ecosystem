import type { NextPage } from "next";
import { useEffect, useState } from "react";
import useBroadcastClient from "../helpers/useBroadcastClient";
import {
  BroadcastQuery,
  BulkBroadcasts,
  Broadcast,
} from "./../proto/operations_ecosys_pb";

const Home: NextPage = () => {
  const [broadcasts, setBroadcasts] = useState<Broadcast[]>([]);

  const broadcastClient = useBroadcastClient();

  const getBroadcasts = async () => {
    const query = new BroadcastQuery();
    // var stream = broadcastClient.findBroadcasts(query, {});

    // stream.on("data", (response: BulkBroadcasts) => {
    //   setBroadcasts(response.getBroadcastsList());
    // });
    await broadcastClient.findBroadcasts(
      query,
      {},
      (err: Error, response: BulkBroadcasts) => {
        if (err) {
          console.log(err);
        } else {
          setBroadcasts(response.getBroadcastsList());
          console.log(response.getBroadcastsList());
        }
      }
    );
  };

  useEffect(() => {
    getBroadcasts();
  }, []);

  return (
    <div>
      <h1>Testing</h1>
      <ul>
        {broadcasts.map((broadcast) => (
          <li key={broadcast.getBroadcastId()}>{broadcast.getContent()}</li>
        ))}
      </ul>
    </div>
  );
};

export default Home;
