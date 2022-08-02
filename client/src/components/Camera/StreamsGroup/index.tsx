import { Group } from "@mantine/core";
import React from "react";
import { useCameraIot } from "../../../helpers/useCameraIotClient";
import StreamContainer from "../StreamContainer";

export default function StreamsGroup() {
  const { cameras } = useCameraIot();

  return (
    <Group>
      {cameras.map((camera) => (
        <StreamContainer
          key={camera.cameraIotId}
          id={camera.cameraIotId}
          name={camera.name}
          videoSrc={camera.camera?.url}
          lightStatus={camera.fireAlarm?.state}
          cpuTemp={camera.cpuTemperature?.temp}
          gateStatus={camera.gate?.state}
        />
      ))}
    </Group>
  );
}
