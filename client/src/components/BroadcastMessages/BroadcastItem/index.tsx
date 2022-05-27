import { createStyles, Text } from "@mantine/core";
import React from "react";

const useStyles = createStyles((theme) => ({
  broadcast: {
    borderRadius: theme.spacing.lg,
    padding: theme.spacing.xs,
    position: "relative",
    paddingRight: "60px",
  },
}));

interface BroadcastItemProps {
  id: number;
  type: string;
  content: string;
  time: string;
  urgency?: string;
}

export default function BroadcastItem(props: BroadcastItemProps) {
  const { classes, theme } = useStyles();
  return (
    <div
      key={props.id}
      className={classes.broadcast}
      style={{
        background:
          props.type == "broadcast"
            ? theme.fn.linearGradient(
                30,
                theme.colors.blue[4],
                theme.colors.blue[3]
              )
            : theme.colors.gray[5],
        color: theme.white,
        alignSelf: props.type == "broadcast" ? "flex-end" : "flex-start",
      }}
    >
      <Text>{props.content}</Text>
      <Text
        size="xs"
        sx={{
          position: "absolute",
          bottom: "4px",
          right: theme.spacing.lg,
        }}
      >
        {props.time}
      </Text>
    </div>
  );
}
