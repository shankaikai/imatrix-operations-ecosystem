import { Stack, Group } from "@mantine/core";
import type { NextPage } from "next";
import ReportContainer from "../components/Reporting/ReportContainer";
import ReportsBar from "../components/Reporting/ReportsBar";
import ReportsList from "../components/Reporting/ReportsList";

const Reporting: NextPage = () => {
  return (
    <Stack>
      <ReportsBar />
      <Group align="flex-start">
        <ReportsList />
        <ReportContainer />
      </Group>
    </Stack>
  );
};

export default Reporting;
