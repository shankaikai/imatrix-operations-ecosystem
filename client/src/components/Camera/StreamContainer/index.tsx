import {
  Card,
  Stack,
  Text,
  Group,
  Box,
  Switch,
  Button,
  UnstyledButton,
} from "@mantine/core";
import React from "react";
import { activateGateSwitch } from "../../../helpers/useCameraIotClient";
import {
  FireAlarmState,
  Gate,
  GateState,
} from "../../../proto/iot_prototype_pb";
import Stream from "../Stream";

interface StreamContainerProps {
  id: any;
  name: string;
  videoSrc?: string;
  lightStatus?: FireAlarmState.AlarmState;
  cpuTemp?: number;
  gateStatus?: GateState.GatePosition;
}

export default function StreamContainer({
  id,
  name,
  videoSrc,
  lightStatus,
  cpuTemp,
  gateStatus,
}: StreamContainerProps) {
  const handleGateSwitch = () => {
    console.log("Activating switch", id);
    gateStatus && activateGateSwitch(id, gateStatus);
  };

  return (
    <Group
      sx={{
        width: "calc(33% - 9px)",
      }}
    >
      <Card sx={{ display: "flex", width: "100%" }} shadow="lg" p="xs">
        <Stack align="center" spacing="xs">
          <Text weight={500}>{name}</Text>
          {videoSrc && <Stream src={videoSrc} id={id} />}
          <div
            style={{
              position: "relative",
              display: "flex",
              flexDirection: "row",
              width: "100%",
              justifyContent: "center",
            }}
          >
            <Text
              sx={{ position: "absolute", left: 0 }}
            >{`CPU: ${cpuTemp}`}</Text>
            <div
              style={{
                height: "24px",
                width: "24px",
                borderRadius: "24px",
                backgroundColor:
                  lightStatus === FireAlarmState.AlarmState.ON
                    ? "red"
                    : "black",
                boxShadow:
                  "0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19)",
              }}
            />
            <UnstyledButton
              sx={{
                right: 0,
                position: "absolute",
              }}
              onChange={handleGateSwitch}
            >
              <Switch
                aria-label="switch"
                color="green"
                checked={gateStatus === GateState.GatePosition.OPEN}
              />
            </UnstyledButton>
          </div>
        </Stack>
      </Card>
    </Group>
  );
}
