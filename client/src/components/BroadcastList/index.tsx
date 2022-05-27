import React, { Dispatch, useState } from "react";
import {
  createStyles,
  Tabs,
  Input,
  Select,
  Text,
  Group,
  TextInput,
  Stack,
  Card,
  Avatar,
  UnstyledButton,
} from "@mantine/core";
import { Photo, MessageCircle, Settings } from "tabler-icons-react";
import BroadcastCard from "./BroadcastCard";
import BroadcastFilter from "./BroadcastFilter";

const useStyles = createStyles((theme) => ({
  tabcontainer: {
    height: "100vh",
    width: "360px",
    padding: theme.spacing.xs,
  },
  filters: {
    columnGap: theme.spacing.xs,
  },
}));

const mockdata = [
  {
    id: 1,
    imgUrl:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQKXW9nkZ9NMKDopl5m6XfT2jIAJamJWko3VpstBIlFDKy4VTVYOx4TVeZ8SpzD1ZGdlAs&usqp=CAU",
    name: "Test Name",
    message: "This is a test message!!!",
    unread: 0,
  },
  {
    id: 2,
    imgUrl:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQKXW9nkZ9NMKDopl5m6XfT2jIAJamJWko3VpstBIlFDKy4VTVYOx4TVeZ8SpzD1ZGdlAs&usqp=CAU",
    name: "Test Name",
    message: "This is a longgggggggggggggggg test message!!!",
    unread: 2,
  },
  {
    id: 3,
    imgUrl:
      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQKXW9nkZ9NMKDopl5m6XfT2jIAJamJWko3VpstBIlFDKy4VTVYOx4TVeZ8SpzD1ZGdlAs&usqp=CAU",
    name: "Test Name",
  },
];

interface BroadcastListProps {
  setSelectedCard: Dispatch<number>;
  selectedCard: number;
}

export default function BroadcastList(props: BroadcastListProps) {
  const { classes } = useStyles();

  return (
    <div className={classes.tabcontainer}>
      <Tabs variant="pills" position="apart">
        <Tabs.Tab label="Direct" icon={<Photo size={14} />}>
          <BroadcastFilter />
          <Stack spacing="xs" mt="xs">
            {mockdata.map((item) => (
              <UnstyledButton
                key={item.id}
                onClick={() => {
                  console.log(item.id);
                  props.setSelectedCard(item.id);
                }}
              >
                <BroadcastCard
                  id={item.id}
                  imgUrl={item.imgUrl}
                  name={item.name}
                  message={item.message}
                  unread={item.unread}
                  selected={props.selectedCard}
                />
              </UnstyledButton>
            ))}
          </Stack>
        </Tabs.Tab>
        <Tabs.Tab label="Groups" icon={<MessageCircle size={14} />}>
          Messages tab content
        </Tabs.Tab>
        <Tabs.Tab label="Forum" icon={<Settings size={14} />}>
          Settings tab content
        </Tabs.Tab>
      </Tabs>
    </div>
  );
}
