import {
  ActionIcon,
  Card,
  Checkbox,
  CheckboxGroup,
  Group,
  Space,
  Stack,
  Text,
} from "@mantine/core";
import dayjs from "dayjs";
import React, { useState } from "react";
import { IoSave } from "react-icons/io5";
import { useReporting } from "../../../helpers/useReportingClient";

export default function ReportContainer() {
  const { selectedReport, reports } = useReporting();

  const handleSave = () => {
    console.log("handleSave called");
  };

  return (
    <Card
      sx={{
        flex: "1",
        display: selectedReport?.incidentReportId ? "default" : "none",
      }}
    >
      <Stack>
        <Group position="apart">
          <Text size="lg" weight={500}>
            {selectedReport?.incidentReportContent?.title}
          </Text>
          <ActionIcon onClick={handleSave}>
            <IoSave />
          </ActionIcon>
        </Group>
        <Space />
        <Stack spacing={0}>
          <Text size="xs">{`Name: ${selectedReport?.creator?.name}`}</Text>
          <Text size="xs">{`Reported on: ${dayjs(
            selectedReport?.creationDate,
            "YYYY-MM-DD HH:mm:ss"
          ).format("D/M/YY [at] HH:mm")}`}</Text>
          <Text size="xs">{`Last updated: ${dayjs(
            selectedReport?.lastModifiedDate,
            "YYYY-MM-DD HH:mm:ss"
          ).format("D/M/YY [at] HH:mm")}`}</Text>
          <Text size="xs">{`Address: ${selectedReport?.incidentReportContent?.address}`}</Text>
        </Stack>
        <Space />
        <Text size="xs">
          {selectedReport?.incidentReportContent?.description}
        </Text>
        <Space />
        <Stack>
          <Group>
            <Text size="xs">Was police notified?</Text>
            <Checkbox
              size="xs"
              checked={selectedReport?.incidentReportContent?.isPoliceNotified}
            />
          </Group>
          <Group>
            <Text size="xs">Was anything stolen?</Text>
            <Checkbox
              size="xs"
              checked={selectedReport?.incidentReportContent?.hasStolenItem}
            />
          </Group>
        </Stack>
      </Stack>
    </Card>
  );
}
