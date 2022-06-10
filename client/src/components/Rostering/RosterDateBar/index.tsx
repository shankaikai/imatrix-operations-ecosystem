import { ActionIcon, Button, Group } from "@mantine/core";
import dayjs from "dayjs";
import React from "react";
import { ChevronLeft, ChevronRight, Send } from "tabler-icons-react";
import getOverallRosterStatus from "../../../helpers/getOverallRosterStatus";
import {
  submitNewRoster,
  submitUpdateRoster,
  useRostering,
} from "../../../helpers/useRosteringClient";
import { Roster } from "../../../proto/operations_ecosys_pb";

export default function RosterDateBar() {
  const {
    rosterDates,
    setOffset,
    offset,
    selectedDate,
    setSelectedDate,
    guardsAssigned,
    rosterBaskets,
    publishDisabled,
  } = useRostering();

  const handleLeftClick = () => {
    setOffset && setOffset(offset - 6);
  };

  const handleRightClick = () => {
    setOffset && setOffset(offset + 6);
  };

  const handleDateClick = (date: Date) => {
    setSelectedDate && setSelectedDate(date);
  };

  const handlePublish = () => {
    console.log("handlePublish called");
    if (getOverallRosterStatus(rosterBaskets) === Roster.Status.REJECTED) {
      selectedDate && submitUpdateRoster(guardsAssigned, selectedDate);
    } else {
      selectedDate && submitNewRoster(guardsAssigned, selectedDate);
    }
  };

  return (
    <Group position="apart">
      <ActionIcon onClick={handleLeftClick}>
        <ChevronLeft color="gray" />
      </ActionIcon>
      {rosterDates.map((date) => (
        <Button
          key={date.getDate()}
          variant={
            date.getDate() === selectedDate?.getDate() ? "filled" : "light"
          }
          color="gray"
          onClick={() => handleDateClick(date)}
        >
          {dayjs(date).format("D MMM")}
        </Button>
      ))}
      <ActionIcon onClick={handleRightClick}>
        <ChevronRight color="gray" />
      </ActionIcon>
      <Button
        leftIcon={<Send size={14} />}
        radius="xl"
        size="xs"
        onClick={handlePublish}
        disabled={publishDisabled}
      >
        Publish
      </Button>
    </Group>
  );
}
