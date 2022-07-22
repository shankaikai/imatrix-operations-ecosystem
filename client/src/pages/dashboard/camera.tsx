import { Stack } from "@mantine/core";
import type { NextPage } from "next";
import Layout from "../../components/Layout/Layout";
import CameraBar from "../../components/Camera/CameraBar";
import StreamsGroup from "../../components/Camera/StreamsGroup";
import { CameraIotProvider } from "../../helpers/useCameraIotClient";

const Camera: NextPage = () => {
  return (
    <Layout>
      <CameraIotProvider>
        <Stack>
          <CameraBar />
          <StreamsGroup />
        </Stack>
      </CameraIotProvider>
    </Layout>
  );
};

export default Camera;
