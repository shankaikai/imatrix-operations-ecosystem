import React from "react";
import {
  Camera,
  Message2,
  LayoutDashboard,
  Report,
  Forms,
} from "tabler-icons-react";
import { ThemeIcon, UnstyledButton, Group, Text } from "@mantine/core";
import Link from "next/link";
import { useRouter } from "next/router";

interface MainLinkProps {
  icon: React.ReactNode;
  color: string;
  label: string;
  link: string;
  selected: boolean;
}

function MainLink({ icon, color, label, link, selected }: MainLinkProps) {
  return (
    <Link href={link}>
      <UnstyledButton
        sx={(theme) => ({
          display: "block",
          width: "100%",
          padding: theme.spacing.xs,
          borderRadius: theme.radius.sm,
          color:
            theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.black,

          backgroundColor: selected
            ? theme.colorScheme === "dark"
              ? theme.colors.dark[6]
              : theme.colors.gray[0]
            : "default",

          "&:hover": {
            backgroundColor:
              theme.colorScheme === "dark"
                ? theme.colors.dark[6]
                : theme.colors.gray[0],
          },
        })}
      >
        <Group>
          <ThemeIcon color={color} variant="light">
            {icon}
          </ThemeIcon>
          <Text size="sm">{label}</Text>
        </Group>
      </UnstyledButton>
    </Link>
  );
}

const links = [
  {
    icon: <LayoutDashboard size={16} />,
    color: "teal",
    label: "Dashboard",
    link: "/",
  },
  {
    icon: <Message2 size={16} />,
    color: "blue",
    label: "Broadcast",
    link: "/broadcasting",
  },
  {
    icon: <Forms size={16} />,
    color: "green",
    label: "Rostering",
    link: "/rostering",
  },
  {
    icon: <Report size={16} />,
    color: "violet",
    label: "Reports",
    link: "/reporting",
  },
  {
    icon: <Camera size={16} />,
    color: "grape",
    label: "Camera",
    link: "/camera",
  },
];

export function MainLinks() {
  const router = useRouter();

  return (
    <div>
      {links.map((link) => (
        <MainLink
          {...link}
          key={link.label}
          selected={router.asPath === link.link}
        />
      ))}
    </div>
  );
}
