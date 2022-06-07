import { Avatar, Stack, Text } from "@mantine/core";
import React from "react";
import { User } from "../../../proto/operations_ecosys_pb";

interface RosterGuardProps {
  guard: User.AsObject;
  withLabels?: boolean;
}

export default function RosterGuard({ guard, withLabels }: RosterGuardProps) {
  return (
    <Stack align="center">
      <Avatar size={80} src={guard.userSecurityImg} radius={100} />
      {withLabels && (
        <Stack align="center">
          <Text>{guard.name}</Text>
          <Text>{guard.phoneNumber}</Text>
        </Stack>
      )}
    </Stack>
  );
}
