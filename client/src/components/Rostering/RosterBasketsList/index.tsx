import { Stack } from "@mantine/core";
import React from "react";
import { useRostering } from "../../../helpers/useRosteringClient";
import RosterBasket from "../RosterBasket";

export default function RosterBasketsLists() {
  const { rosterBaskets } = useRostering();

  return (
    <Stack>
      {rosterBaskets?.map((basket, index) => (
        <RosterBasket key={basket.aifsId} basket={basket} index={index} />
      ))}
    </Stack>
  );
}
