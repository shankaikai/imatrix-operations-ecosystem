import { Stack } from "@mantine/core";
import React from "react";
import ReportSmallCard from "../ReportSmallCard";

export default function ReportsList() {
  return (
    <Stack
      sx={{
        width: "30%",
      }}
    >
      <ReportSmallCard />
      <ReportSmallCard />
      <ReportSmallCard />
      <ReportSmallCard />
      <ReportSmallCard />
      <ReportSmallCard />
    </Stack>
  );
}
