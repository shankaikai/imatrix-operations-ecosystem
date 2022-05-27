import React from "react";
import {
  Navbar,
  Group,
  Code,
  ScrollArea,
  createStyles,
  Text,
} from "@mantine/core";
import {
  Gauge,
  PresentationAnalytics,
  FileAnalytics,
  Adjustments,
} from "tabler-icons-react";
import { UserButton } from "../UserButton";
import { Links } from "../NavbarLinksGroups";
// import { Logo } from './Logo';

const mockdata = [
  { label: "Broadcast", icon: Gauge, link: "/broadcasting" },
  { label: "Rostering", icon: PresentationAnalytics, link: "/rostering" },
  { label: "Camera", icon: FileAnalytics, link: "/camera" },
  { label: "Settings", icon: Adjustments, link: "/settings" },
];

const useStyles = createStyles((theme) => ({
  navbar: {
    backgroundColor:
      theme.colorScheme === "dark" ? theme.colors.dark[6] : theme.white,
    paddingBottom: 0,
    height: "vh",
    width: "380px",
  },

  header: {
    padding: theme.spacing.md,
    color: theme.colorScheme === "dark" ? theme.white : theme.black,
    borderBottom: `1px solid ${
      theme.colorScheme === "dark" ? theme.colors.dark[4] : theme.colors.gray[3]
    }`,
  },

  links: {
    marginLeft: -theme.spacing.md,
    marginRight: -theme.spacing.md,
  },

  linksInner: {
    paddingLeft: theme.spacing.md,
    paddingRight: theme.spacing.md,
  },

  footer: {
    borderTop: `1px solid ${
      theme.colorScheme === "dark" ? theme.colors.dark[4] : theme.colors.gray[3]
    }`,
    paddingLeft: theme.spacing.xs,
    paddingRight: theme.spacing.xs,
  },
}));

export function Sidebar() {
  const { classes } = useStyles();
  const links = mockdata.map((item) => <Links {...item} key={item.label} />);

  return (
    <Navbar className={classes.navbar}>
      <Navbar.Section className={classes.header}>
        <Group position="apart">
          {/* <Logo width={120} /> */}
          <Text size="xl">iMatrix</Text>
        </Group>
      </Navbar.Section>

      <Navbar.Section grow className={classes.links} component={ScrollArea}>
        <div className={classes.linksInner}>{links}</div>
      </Navbar.Section>

      <Navbar.Section className={classes.footer}>
        <UserButton
          image="https://images.unsplash.com/photo-1508214751196-bcfd4ca60f91?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=255&q=80"
          name="Ann Nullpointer"
          email="anullpointer@yahoo.com"
        />
      </Navbar.Section>
    </Navbar>
  );
}
