import React from "react";
import {
  AppShell,
  Navbar,
  Header,
  Group,
  ActionIcon,
  useMantineColorScheme,
  Text,
  Indicator,
} from "@mantine/core";
import { Sun, MoonStars, Bell } from "tabler-icons-react";

export default function MainHeader() {
  const { colorScheme, toggleColorScheme } = useMantineColorScheme();

  const handleBellClick = () => {
    // TODO: Code out notifcations drop down logic
  };

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
        <Group>
          <ActionIcon onClick={handleBellClick}>
            <Indicator color="red" size={10} offset={2} withBorder>
              <Bell size={16} />
            </Indicator>
          </ActionIcon>
          <ActionIcon
            variant="default"
            onClick={() => toggleColorScheme()}
            size={30}
          >
            {colorScheme === "dark" ? (
              <Sun size={16} />
            ) : (
              <MoonStars size={16} />
            )}
          </ActionIcon>
        </Group>
      </Group>
    </Header>
  );
}
