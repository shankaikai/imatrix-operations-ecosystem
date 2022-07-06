import { Avatar, Card, Group, Stack, Text } from "@mantine/core";
import dayjs from "dayjs";
import React from "react";
import { useReporting } from "../../../helpers/useReportingClient";
import { IncidentReport } from "../../../proto/operations_ecosys_pb";

interface ReportSmallCardProps {
  id: number;
  title: string;
  sender: string;
  senderImg: string;
  aifsId?: number;
  creationDate: string;
  updateDate: string;
  selected?: boolean;
  report?:IncidentReport.AsObject
}

export default function ReportSmallCard({
  selected,
  id,
  title,
  sender,
  senderImg,
  aifsId,
  creationDate,
  updateDate,
  report,
}: ReportSmallCardProps) {
  const { setSelectedReport } = useReporting();

  const handleClick = () => {
    console.log("set id to ", id);
    setSelectedReport && report && setSelectedReport(report);
  };

  return (
    <Card
      shadow="sm"
      p="xs"
      sx={(theme) => ({
        backgroundColor: selected
          ? theme.colorScheme === "dark"
            ? theme.colors.dark[4]
            : theme.colors.gray[2]
          : "default",

        "&:hover": {
          backgroundColor:
            theme.colorScheme === "dark"
              ? theme.colors.dark[5]
              : theme.colors.gray[1],
        },
      })}
      onClick={handleClick}
    >
      <Stack>
        <Text weight={500}>{title}</Text>
        <Group spacing="xs">
          <Avatar radius="xl" src={senderImg}/>
          <Text size="xs" color="dimmed">
            {sender}
          </Text>
          <Text>•</Text>
          <Text size="xs" color="dimmed">
            {`Updated on ${dayjs(updateDate, "YYYY-MM-DD HH:mm:ss").format(
              "D/M/YY [at] HH:mm"
            )}`}
          </Text>
        </Group>
      </Stack>
    </Card>
  );
}
