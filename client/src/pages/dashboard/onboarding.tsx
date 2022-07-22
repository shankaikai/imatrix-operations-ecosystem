import { Stack, Group } from "@mantine/core";
import type { NextPage } from "next";
import Layout from "../../components/Layout/Layout";
import ReportContainer from "../../components/Reporting/ReportContainer";
import ReportsBar from "../../components/Reporting/ReportsBar";
import ReportsList from "../../components/Reporting/ReportsList";
import ReportsModal from "../../components/Reporting/ReportsModal";
import { ReportingProvider } from "../../helpers/useReportingClient";

const Reporting: NextPage = () => {
  return (
    <Layout>
      <ReportingProvider>
        <Stack>
          <ReportsBar />
          <Group align="flex-start">
            <ReportsList />
            <ReportContainer />
            <ReportsModal />
          </Group>
        </Stack>
      </ReportingProvider>
    </Layout>
  );
};

export default Reporting;
