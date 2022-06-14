import {
  Button,
  Group,
  MultiSelect,
  Stack,
  Textarea,
  Text,
  Badge,
} from "@mantine/core";
import React, { Dispatch, useState } from "react";
import { Send } from "tabler-icons-react";
import { submitNewBroadcast } from "../../../helpers/useBroadcastClient";
import CustomRecipientsBadge from "./CustomRecipientsBadge";
import CustomUrgencyBadge from "./CustomUrgencyBadge";

interface NewBroadcastProps {
  setModelOpen: Dispatch<boolean>;
}

const recipients = [
  { label: "All", value: "all" },
  { label: "AIFS 1", value: "1" },
  { label: "AIFS 2", value: "2" },
  { label: "AIFS 3", value: "3" },
];

const urgencys = ["Urgent", "Not Urgent"];

export default function NewBroadcast({ setModelOpen }: NewBroadcastProps) {
  const [recipient, setRecipient] = useState<string[]>([]);
  const [urgency, setUrgency] = useState<string>("");
  const [message, setMessage] = useState<string>("");

  const handleSubmit = async () => {
    console.log("Submitting new broadcast");

    await submitNewBroadcast({
      recipient,
      urgency,
      message,
    });

    // TODO: Loading overlay before closing modal
    setModelOpen(false);
  };

  return (
    <form>
      <Stack>
        <Group>
          <Text size="sm">Recipients:</Text>
          {recipients.map(({ label, value }) => (
            <CustomRecipientsBadge
              key={value}
              label={label}
              value={value}
              setValue={setRecipient}
              active={recipient.includes(value)}
            />
          ))}
        </Group>
        <Group>
          <Text size="sm">Urgency:</Text>
          {urgencys.map((_urgency) => (
            <CustomUrgencyBadge
              key={_urgency}
              label={_urgency}
              setValue={setUrgency}
              active={_urgency === urgency}
            />
          ))}
        </Group>
        <Textarea
          data-testid="newBroadcastMessageInput"
          label="Message"
          placeholder="Enter your broadcast message"
          autosize
          minRows={2}
          maxRows={5}
          value={message}
          onChange={(event) => setMessage(event.currentTarget.value)}
        />
        <Group position="center">
          <Button
            data-testid="submitBroadcast"
            leftIcon={<Send size={16} />}
            onClick={handleSubmit}
          >
            Submit
          </Button>
        </Group>
      </Stack>
    </form>
  );
}
