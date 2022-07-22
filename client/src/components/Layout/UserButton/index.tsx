import React, { useState } from "react";
import {
  UnstyledButton,
  UnstyledButtonProps,
  Group,
  Avatar,
  Text,
  createStyles,
  Divider,
  Modal,
  Button,
} from "@mantine/core";
import { ChevronRight } from "tabler-icons-react";
import { signOut } from "../../../helpers/useUserProvider";
import { useRouter } from "next/router";

const useStyles = createStyles((theme) => ({
  user: {
    display: "block",
    width: "100%",
    padding: theme.spacing.md,
    color: theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.black,

    "&:hover": {
      backgroundColor:
        theme.colorScheme === "dark"
          ? theme.colors.dark[8]
          : theme.colors.gray[0],
    },
  },
}));

//@ts-ignore
interface UserButtonProps extends UnstyledButtonProps {
  image: string;
  name: string;
  email: string;
  icon?: React.ReactNode;
}

export function UserButton({
  image,
  name,
  email,
  icon,
  ...others
}: UserButtonProps) {
  const { classes } = useStyles();

  const [modalOpen, setModalOpen] = useState(false);
  const router = useRouter();

  const logOut = () => {
    signOut();
    router.push("/login");
  };
  return (
    <div>
      <Modal
        opened={modalOpen}
        onClose={() => setModalOpen(false)}
        title="Sign out"
      >
        <Text size="md" weight={500}>
          Are you sure you want to sign out?
        </Text>
        <Group sx={{ width: "100%", justifyContent: "end", marginTop: "32px" }}>
          <Button onClick={logOut}>Sign out</Button>
          <Button onClick={() => setModalOpen(false)} variant="outline">
            Cancel
          </Button>
        </Group>
      </Modal>
      <Divider />
      <UnstyledButton
        className={classes.user}
        onClick={() => setModalOpen(true)}
        {...others}
      >
        <Group>
          <Avatar src={image} radius="xl" />

          <div style={{ flex: 1 }}>
            <Text size="sm" weight={500}>
              {name}
            </Text>

            <Text color="dimmed" size="xs">
              {email}
            </Text>
          </div>

          {icon || <ChevronRight size={14} />}
        </Group>
      </UnstyledButton>
    </div>
  );
}
