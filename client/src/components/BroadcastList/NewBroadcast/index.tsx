import {
  Button,
  Group,
  MultiSelect,
  Stack, Textarea
} from "@mantine/core";
import React, { useState } from "react";
import { Send } from "tabler-icons-react";

export default function NewBroadcast() {
  const [recipient, setRecipient] = useState<string[]>(["All"]);
  const [urgency, setUrgency] = useState<string[]>(["All"]);

  return (
    <form>
      <Stack>
        <MultiSelect
          data={["All", "AIFS 1", "AIFS 2", "AIFS 3"]}
          label="Recipients"
          defaultValue={["All"]}
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
        />
        <Textarea
          label="Message"
          placeholder="Enter your broadcast message"
          autosize
          minRows={2}
          maxRows={5}
        />
        <Group position="center">
          <Button leftIcon={<Send size={16} />} >
            Submit
          </Button>
        </Group>
      </Stack>
    </form>
  );
}
