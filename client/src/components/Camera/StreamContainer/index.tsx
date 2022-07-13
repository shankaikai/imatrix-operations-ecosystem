import { Card, Stack, Text, Group, Box, Switch } from "@mantine/core";
import React from "react";
import Stream from "../Stream";

interface StreamContainerProps {
  id: any;
  videoSrc: string;
  lightStatus: boolean;
  cpuTemp: number;
}

export default function StreamContainer({
  id,
  videoSrc,
  lightStatus,
  cpuTemp,
}: StreamContainerProps) {
  return (
    <Group
      sx={{
        width: "calc(33% - 9px)",
      }}
    >
      <Card sx={{ display: "flex" }} shadow="lg" p="xs">
        <Stack align="center" spacing="xs">
          <Text weight={500}>S8</Text>
          <Stream src={videoSrc} id={id} />
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
                backgroundColor: lightStatus ? "red" : "black",
                boxShadow:
                  "0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19)",
              }}
            />
            <Switch
              sx={{
                right: 0,
                position: "absolute",
              }}
              color="green"
            />
          </div>
        </Stack>
      </Card>
    </Group>
  );
}
