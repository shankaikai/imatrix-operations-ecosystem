import { ComponentMeta, ComponentStory } from "@storybook/react";
import StreamContainer from "../../components/Camera/StreamContainer";
import { FireAlarmState, GateState } from "../../proto/iot_prototype_pb";

export default {
  title: "Camera/StreamContainer",
  component: StreamContainer,
} as ComponentMeta<typeof StreamContainer>;

const Template: ComponentStory<typeof StreamContainer> = (args) => (
  <StreamContainer {...args} />
);

export const Default = Template.bind({});
Default.args = {
  id: 1,
  name: "C1",
  videoSrc:
    "http://aifs.lunarcloud.org:8080/beabf5988b310b8a6e51fec9a2d3f4f8/hls/Xt5RWwa66W/kexVHDZPK7/s.m3u8",
  lightStatus: FireAlarmState.AlarmState.ON,
  cpuTemp: 45,
  gateStatus: GateState.GatePosition.CLOSED,
};
