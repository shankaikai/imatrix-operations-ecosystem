import { Avatar, Card, Group, Stack, Text } from "@mantine/core";
import React from "react";
import { useReporting } from "../../../helpers/useReportingClient";

interface ReportSmallCardProps {
  id?: number;
  title?: string;
  sender?: string;
  aifsId?: number;
  updateDate?: Date;
  selected?: boolean;
}

export default function ReportSmallCard({ selected }: ReportSmallCardProps) {
  const { setSelectedReport } = useReporting();

  const handleClick = () => {
    console.log("ReportSmallCard clicked");
    setSelectedReport && setSelectedReport(1);
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
        <Text weight={500}>Camera Faulty.doc</Text>
        <Group spacing="xs">
          <Avatar radius="xl" />
          <Text size="xs" color="dimmed">
            Philip Wee (AIFS1)
          </Text>
          <Text>•</Text>
          <Text size="xs" color="dimmed">
            Updated on 7/6/22 at 19:04
          </Text>
        </Group>
      </Stack>
    </Card>
  );
}
