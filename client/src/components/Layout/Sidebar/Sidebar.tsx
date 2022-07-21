import { Navbar } from "@mantine/core";
import React from "react";
import { MainLinks } from "../MainLinks/MainLinks";
import { UserButton } from "../UserButton";

export default function Sidebar() {
  return (
    <Navbar width={{ base: 280 }} height={"calc(100vh - 60px)"} p="xs">
      <Navbar.Section grow mt="xs">
        <MainLinks />
      </Navbar.Section>
      <Navbar.Section>
        {/* TODO: Change this to accept user data from context */}
        <UserButton
          image="https://images.unsplash.com/photo-1508214751196-bcfd4ca60f91?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=255&q=80"
          name="Ann Nullpointer"
          email="anullpointer@yahoo.com"
        />
      </Navbar.Section>
    </Navbar>
  );
}
