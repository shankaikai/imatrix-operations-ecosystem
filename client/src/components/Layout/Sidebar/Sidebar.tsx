import { Navbar } from "@mantine/core";
import { useRouter } from "next/router";
import React from "react";
import { useUserProvider } from "../../../helpers/useUserProvider";
import { MainLinks } from "../MainLinks/MainLinks";
import { UserButton } from "../UserButton";

export default function Sidebar() {
  const { name, email, image } = useUserProvider();
  const router = useRouter();
  return (
    <Navbar width={{ base: 280 }} height={"calc(100vh - 60px)"} p="xs">
      <Navbar.Section grow mt="xs">
        <MainLinks />
      </Navbar.Section>
      <Navbar.Section>
        {/* TODO: Change this to accept user data from context */}
        <UserButton image={image} name={name} email={email} />
      </Navbar.Section>
    </Navbar>
  );
}
