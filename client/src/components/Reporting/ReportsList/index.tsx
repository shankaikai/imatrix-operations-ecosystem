import { ScrollArea, Stack } from "@mantine/core";
import React from "react";
import ReportSmallCard from "../ReportSmallCard";

export default function ReportsList() {
  return (
    <ScrollArea
      sx={{
        width: "35%",
        height: "calc(100vh - 132px)",
      }}
    >
      <Stack>
        <ReportSmallCard selected/>
        <ReportSmallCard />
        <ReportSmallCard />
        <ReportSmallCard />
        <ReportSmallCard />
        <ReportSmallCard />
      </Stack>
    </ScrollArea>
  );
}
