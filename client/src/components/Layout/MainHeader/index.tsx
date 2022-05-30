import React from "react";
import {
  AppShell,
  Navbar,
  Header,
  Group,
  ActionIcon,
  useMantineColorScheme,
  Text,
} from "@mantine/core";
import { Sun, MoonStars } from "tabler-icons-react";

export default function MainHeader() {
  const { colorScheme, toggleColorScheme } = useMantineColorScheme();

  return (
    <Header height={60} sx={{ width: "100vw" }} px="sm">
      <Group sx={{ height: "100%", width: "100%" }} position="apart">
        <Text
          sx={{
            width: 280,
          }}
        >
          iMatrix Operations Dashboard
        </Text>

        <ActionIcon
          variant="default"
          onClick={() => toggleColorScheme()}
          size={30}
        >
          {colorScheme === "dark" ? <Sun size={16} /> : <MoonStars size={16} />}
        </ActionIcon>
      </Group>
    </Header>
  );
}
