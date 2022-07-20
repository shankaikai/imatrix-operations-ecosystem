import { Group } from "@mantine/core";
import React from "react";
import { useCameraIot } from "../../../helpers/useCameraIotClient";
import StreamContainer from "../StreamContainer";

const fakeData = [
  {
    id: 1,
    videoSrc:
      "http://aifs.lunarcloud.org:8080/beabf5988b310b8a6e51fec9a2d3f4f8/hls/Xt5RWwa66W/kexVHDZPK7/s.m3u8",
    lightStatus: true,
    cpuTemp: 56,
  },
  {
    id: 2,
    videoSrc:
      "https://devstreaming-cdn.apple.com/videos/streaming/examples/img_bipbop_adv_example_fmp4/master.m3u8",
    lightStatus: false,
    cpuTemp: 45,
  },
  {
    id: 3,
    videoSrc:
      "https://devstreaming-cdn.apple.com/videos/streaming/examples/img_bipbop_adv_example_fmp4/master.m3u8",
    lightStatus: true,
    cpuTemp: 56,
  },
];

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
        />
      ))}
    </Group>
  );
}
