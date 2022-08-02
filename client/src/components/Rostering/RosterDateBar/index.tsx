import { ActionIcon, Button, Group } from "@mantine/core";
import dayjs from "dayjs";
import React, { useEffect } from "react";
import { ChevronLeft, ChevronRight, Refresh, Send } from "tabler-icons-react";
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
    setPublishDisabled,
    refreshState,
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
      selectedDate &&
        setPublishDisabled &&
        submitUpdateRoster(guardsAssigned, selectedDate, setPublishDisabled, rosterBaskets);
      }
    if (getOverallRosterStatus(rosterBaskets) === Roster.Status.IS_DEFAULT) {
      selectedDate &&
        setPublishDisabled &&
        submitNewRoster(guardsAssigned, selectedDate, setPublishDisabled);
    }
  
  };

  useEffect(() => {
    console.log("rosterbasketstatus", getOverallRosterStatus(rosterBaskets));
    if (
      getOverallRosterStatus(rosterBaskets) === Roster.Status.PENDING ||
      getOverallRosterStatus(rosterBaskets) === Roster.Status.CONFIRMED
    ) {
      setPublishDisabled && setPublishDisabled(true);
    }
    if (
      getOverallRosterStatus(rosterBaskets) === Roster.Status.IS_DEFAULT ||
      getOverallRosterStatus(rosterBaskets) === Roster.Status.REJECTED
    ) {
      setPublishDisabled && setPublishDisabled(false);
    }
  }, [rosterBaskets]);

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
      <ActionIcon onClick={()=>refreshState && refreshState()}>
          <Refresh size={16} />
      </ActionIcon>
    </Group>
  );
}
