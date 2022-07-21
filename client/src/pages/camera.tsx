<<<<<<< HEAD
import type { NextPage } from "next";

const Camera: NextPage = () => {
  return (
    <div>
      <h1>Camera</h1>
    </div>
  );
};

export default Camera;
=======
import { Stack } from "@mantine/core";
import type { NextPage } from "next";
import CameraBar from "../components/Camera/CameraBar";
import StreamsGroup from "../components/Camera/StreamsGroup";
import { CameraIotProvider } from "../helpers/useCameraIotClient";

const Camera: NextPage = () => {
  return (
    <CameraIotProvider>
      <Stack>
        <CameraBar />
        <StreamsGroup />
      </Stack>
    </CameraIotProvider>
  );
};

export default Camera;
>>>>>>> feat/camera_system
