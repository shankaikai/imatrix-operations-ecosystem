import { Group, Button, ActionIcon } from "@mantine/core";
import dayjs from "dayjs";
import React, { useState } from "react";
import { ChevronLeft, ChevronRight, Send } from "tabler-icons-react";
import getRosterDates from "../../../helpers/getRosterDates";
import { useRostering } from "../../../helpers/useRosteringClient";

export default function RosterDateBar() {
  const { rosterDates, setOffset, offset, selectedDate, setSelectedDate } =
    useRostering();

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
      >
        Publish
      </Button>
    </Group>
  );
}
