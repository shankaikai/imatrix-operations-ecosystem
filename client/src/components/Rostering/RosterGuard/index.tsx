import { Avatar, Group, Indicator, Popover, Stack, Text } from "@mantine/core";
import React, { useState } from "react";
import { useDrag } from "react-dnd";
import { Roster, User } from "../../../proto/operations_ecosys_pb";

interface RosterGuardProps {
  guard?: User.AsObject;
  withLabels?: boolean;
  index: number;
  nonDraggable?: boolean;
  status?: Roster.Status;
}

export interface DraggableGuard {
  id: number;
  index: number;
}

export default function RosterGuard({
  guard,
  withLabels,
  index,
  nonDraggable,
  status,
}: RosterGuardProps) {
  const [{ isDragging }, drag] = useDrag(() => ({
    type: "guard",
    item: { id: guard?.userId, index },
    collect: (monitor) => {
      return {
        isDragging: monitor.isDragging(),
      };
    },
  }));

  const [opened, setOpened] = useState(false);

  return (
    <div
      ref={nonDraggable ? null : drag}
      style={{
        opacity: isDragging ? 0 : 1,
      }}
    >
      <Popover
        opened={opened}
        position="bottom"
        placement="center"
        withArrow
        trapFocus={false}
        closeOnEscape={false}
        transition="pop-top-left"
        width={260}
        styles={{ body: { pointerEvents: "none" } }}
        target={
          <Group
            onMouseEnter={withLabels ? undefined : () => setOpened(true)}
            onMouseLeave={() => setOpened(false)}
          >
            <Indicator
              offset={10}
              size={14}
              withBorder
              disabled={!withLabels}
              color={
                status === Roster.Status.PENDING
                  ? "orange"
                  : status === Roster.Status.CONFIRMED
                  ? "green"
                  : status === Roster.Status.IS_DEFAULT
                  ? "blue"
                  : "red"
              }
            >
              <Avatar size={80} src={guard?.userSecurityImg} radius={100} />
            </Indicator>
            {withLabels && (
              <Stack align="left">
                <Text weight={600}>{guard?.name}</Text>
                <Text>{guard?.phoneNumber}</Text>
              </Stack>
            )}
          </Group>
        }
      >
        <Text weight={600}>{guard?.name}</Text>
        <Text>{guard?.phoneNumber}</Text>
      </Popover>
    </div>
  );
}
