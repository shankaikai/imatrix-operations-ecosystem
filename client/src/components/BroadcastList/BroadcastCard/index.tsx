import {
  Card,
  Group,
  Avatar,
  Stack,
  Text,
  Space,
  Badge,
  createStyles,
} from "@mantine/core";
import React from "react";
import { Checks } from "tabler-icons-react";
import truncateString from "../../../helpers/truncateString";

interface BroadcastCardProps {
  id: number;
  imgUrl: string;
  name: string;
  message?: string;
  unread?: number;
  selected?: number;
}

const useStyles = createStyles((theme) => ({}));

export default function BroadcastCard(props: BroadcastCardProps) {
  const { theme } = useStyles();

  return (
    <Card
      key={props.name}
      p="sm"
      color="gray"
      shadow="sm"
      sx={(theme) => ({
        backgroundColor:
          props.selected == props.id ? theme.colors.gray[3] : "white",
      })}
    >
      <Group position="apart">
        <Avatar src={props.imgUrl} size="lg" radius="xl" />
        <Stack sx={{ width: "240px" }}>
          <Text size="sm" weight={500}>
            {props.name}
          </Text>
          {props.message ? (
            <Group position="apart" pr="0">
              <Text size="sm" color="dimmed">
                {truncateString(props.message, 25)}
              </Text>
              {props.unread ? (
                <Badge color="red" variant="filled">
                  {props.unread}
                </Badge>
              ) : (
                <Checks color="green" />
              )}
            </Group>
          ) : (
            <Space h="sm" />
          )}
        </Stack>
      </Group>
    </Card>
  );
}
