import React from "react";
import {
  Button,
  Card,
  createStyles,
  Group,
  Input,
  Stack,
  Text,
  TextInput,
  UnstyledButton,
} from "@mantine/core";
import * as dayjs from "dayjs";
import BroadcastItem from "./BroadcastItem";
import { MoodSmile, AlertTriangle, Paperclip, Send } from "tabler-icons-react";

const useStyles = createStyles((theme) => ({
  stack: {
    height: "100vh",
    flexGrow: 1,
    backgroundColor:
      theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.white,
  },
}));

const mockdata = [
  {
    id: 1,
    time: "9:00",
    content: "hello this is message",
    urgency: "high",
    type: "broadcast",
  },
  {
    id: 2,
    time: "9:20",
    content: "hello this is message",
    urgency: "medium",
    type: "broadcast",
  },
  {
    id: 3,
    time: "9:22",
    content: "Acknowledged by AIFS 3",
    type: "acknowledgement",
  },
];

interface BroadcastMessagesProps {
  selectedCard: number;
}

export default function BroadcastMessages(props: BroadcastMessagesProps) {
  const { classes, theme } = useStyles();

  const rightSection = (
    //@ts-ignore
    <Group position="right" spacing="3px">
      <UnstyledButton>
        <MoodSmile color={theme.colors.gray[5]} />
      </UnstyledButton>
      <UnstyledButton>
        <AlertTriangle color={theme.colors.red[5]} />
      </UnstyledButton>
    </Group>
  );

  const attachFile = (
    <UnstyledButton>
      <AlertTriangle color={theme.colors.red[5]} />
    </UnstyledButton>
  );

  return (
    <Stack
      className={classes.stack}
      p="lg"
      justify="flex-end"
      sx={{ display: props.selectedCard == -1 ? "none" : "default" }}
    >
      <Stack spacing="lg" align="flex-start" justify="flex-end">
        {mockdata.map((msg) => (
          <BroadcastItem
            key={msg.id}
            id={msg.id}
            type={msg.type}
            content={msg.content}
            time={msg.time}
            urgency={msg.urgency}
          />
        ))}
      </Stack>
      <Group position="apart">
        <TextInput
          icon={<Paperclip />}
          rightSectionWidth={70}
          rightSection={rightSection}
          variant="filled"
          radius="xl"
          placeholder="Type a message here..."
          sx={{
            flex: "1 1 auto",
          }}
        />
        <Button variant="light" color="gray" radius="xl">
          <Send color={theme.colors.blue[4]} />
        </Button>
      </Group>
    </Stack>
  );
}
