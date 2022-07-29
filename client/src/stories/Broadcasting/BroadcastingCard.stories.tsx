import React from "react";
import { ComponentStory, ComponentMeta } from "@storybook/react";
import {
  AIFSRecipient,
  RecipientDetails,
} from "../../helpers/recipientsFormatter";
import BroadcastCard from "../../components/Broadcasting/BroadcastCard/BroadcastCard";

const mockRecipientDetails: RecipientDetails[] = [
  {
    id: 1,
    phone: "92224444",
    acknowledged: true,
    img: "https://picsum.photos/200",
    name: "Some Name",
  },
  {
    id: 2,
    phone: "92224444",
    acknowledged: false,
    img: "https://picsum.photos/201",
    name: "Some Name",
  },
  {
    id: 3,
    phone: "92224444",
    acknowledged: true,
    img: "https://picsum.photos/202",
    name: "Some Name",
  },
];

const mockRecipientDetails2: RecipientDetails[] = [
  {
    id: 1,
    phone: "92224444",
    acknowledged: true,
    img: "https://picsum.photos/200",
    name: "Some Name",
  },
  {
    id: 2,
    phone: "92224444",
    acknowledged: true,
    img: "https://picsum.photos/201",
    name: "Some Name",
  },
  {
    id: 3,
    phone: "92224444",
    acknowledged: true,
    img: "https://picsum.photos/202",
    name: "Some Name",
  },
];

const mockAifs: AIFSRecipient[] = [
  {
    id: "AIFS 1",
    location: "Location",
    allAcknowledged: false,
    users: mockRecipientDetails,
  },
  {
    id: "AIFS 2",
    location: "Location",
    allAcknowledged: true,
    users: mockRecipientDetails2,
  },
];

const mockAifs2: AIFSRecipient[] = [
  {
    id: "AIFS 1",
    location: "Location",
    allAcknowledged: true,
    users: mockRecipientDetails2,
  },
  {
    id: "AIFS 2",
    location: "Location",
    allAcknowledged: true,
    users: mockRecipientDetails2,
  },
];

export default {
  title: "Broadcasting/BroadcastCard",
  component: BroadcastCard,
} as ComponentMeta<typeof BroadcastCard>;

const Template: ComponentStory<typeof BroadcastCard> = (args) => (
  <BroadcastCard {...args} />
);

export const NotAllAcknowledged = Template.bind({});
NotAllAcknowledged.args = {
  content: "Mock Content",
  date: "2022-10-11 23:59",
  aifs: mockAifs,
};

export const AllAcknowledged = Template.bind({});
AllAcknowledged.args = {
  content: "Mock Content",
  date: "2022-10-11 23:59",
  aifs: mockAifs2,
};
