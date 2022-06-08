import { Button, Group, MultiSelect, Stack, Textarea } from "@mantine/core";
import React, { Dispatch, useState } from "react";
import { Send } from "tabler-icons-react";
import { submitNewBroadcast } from "../../../helpers/useBroadcastClient";

interface NewBroadcastProps {
  setModelOpen: Dispatch<boolean>;
}

export default function NewBroadcast({ setModelOpen }: NewBroadcastProps) {
  const [recipient, setRecipient] = useState<string[]>([]);
  const [urgency, setUrgency] = useState<string[]>([]);
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
        <MultiSelect
          data={[
            {
              label: "All",
              value: "all",
            },
            {
              label: "AIFS 1",
              value: "1",
            },
            {
              label: "AIFS 2",
              value: "2",
            },
            {
              label: "AIFS 3",
              value: "3",
            },
          ]}
          label="Recipients"
          clearButtonLabel="Clear selection"
          clearable
          value={recipient}
          onChange={setRecipient}
        />
        <MultiSelect
          data={["Low", "Medium", "High"]}
          label="Urgency"
          placeholder="Choose your urgency level"
          clearButtonLabel="Clear selection"
          clearable
          maxSelectedValues={1}
          value={urgency}
          onChange={setUrgency}
        />
        <Textarea
          label="Message"
          placeholder="Enter your broadcast message"
          autosize
          minRows={2}
          maxRows={5}
          value={message}
          onChange={(event) => setMessage(event.currentTarget.value)}
        />
        <Group position="center">
          <Button leftIcon={<Send size={16} />} onClick={handleSubmit}>
            Submit
          </Button>
        </Group>
      </Stack>
    </form>
  );
}
