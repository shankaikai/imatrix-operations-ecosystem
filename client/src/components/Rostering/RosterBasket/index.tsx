import { Card, Group, Stack, Text, useMantineTheme } from "@mantine/core";
import dayjs from "dayjs";
import React from "react";
import { useDrop } from "react-dnd";
import addGuardToGuardsAssigned from "../../../helpers/addGuardsToGuardsAssigned";
import { useRostering } from "../../../helpers/useRosteringClient";
import { Roster } from "../../../proto/operations_ecosys_pb";
import RosterGuard, { DraggableGuard } from "../RosterGuard";

interface RosterCardProps {
  basket: Roster.AsObject;
  index: number;
}

export default function RosterBasket({ basket, index }: RosterCardProps) {
  const theme = useMantineTheme();
  const { guardsAssigned, selectedDate, setGuardsAssigned } = useRostering();

  const [{ isOver }, drop] = useDrop({
    accept: "guard",
    drop: (guard: DraggableGuard) =>
      addGuardToGuardsAssigned(
        guard.id,
        guard.index,
        basket.aifsId,
        selectedDate,
        setGuardsAssigned
      ),
    collect: (monitor) => ({
      isOver: monitor.isOver(),
    }),
  });

  return (
    <Card>
      <Stack>
        <Group>
          <Text>{`AIFS ${basket.aifsId}`}</Text>
        </Group>
        <div
          ref={drop}
          style={{
            border: `2px dashed ${
              theme.colorScheme === "dark"
                ? theme.colors.dark[3]
                : theme.colors.gray[5]
            }`,
            borderStyle: "dashed",
            height: "auto",
            padding: theme.spacing.xs,
            borderRadius: theme.spacing.sm,
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            background: isOver ? theme.colors.gray[1] : "default",
          }}
        >
          {selectedDate &&
          guardsAssigned[dayjs(selectedDate).format("YYYY-MM-DD")] ? (
            guardsAssigned[dayjs(selectedDate).format("YYYY-MM-DD")][
              basket.aifsId
            ].map((guard, index) => {
              return (
                <RosterGuard
                  key={guard.employee?.userId}
                  guard={guard.employee}
                  index={index}
                  status={basket.status}
                  withLabels
                />
              );
            })
          ) : (
            <Text>Please drag over a user</Text>
          )}
        </div>
      </Stack>
    </Card>
  );
}
