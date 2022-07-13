import { Stack } from "@mantine/core";
import type { NextPage } from "next";
import CameraBar from "../components/Camera/CameraBar";
import Stream from "../components/Camera/Stream";
import StreamContainer from "../components/Camera/StreamContainer";
import StreamsGroup from "../components/Camera/StreamsGroup";

const Camera: NextPage = () => {
  return (
    <Stack>
      <CameraBar />
      <StreamsGroup />
    </Stack>
  );
};

export default Camera;
