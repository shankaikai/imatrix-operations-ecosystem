import { ComponentMeta, ComponentStory } from "@storybook/react";
import AcknowledgementToolTip from "../../components/Broadcasting/BroadcastCard/AcknowledgementToolTip";
import {
  AIFSRecipient,
  RecipientDetails,
} from "../../helpers/recipientsFormatter";

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

const mockAifRecipient: AIFSRecipient = {
  id: "AIFS 1",
  location: "Location",
  allAcknowledged: false,
  users: mockRecipientDetails,
};

const mockAifRecipient2: AIFSRecipient = {
  id: "AIFS 1",
  location: "Location",
  allAcknowledged: true,
  users: mockRecipientDetails2,
};

export default {
  title: "Broadcasting/AcknowledgementToolTip",
  component: AcknowledgementToolTip,
} as ComponentMeta<typeof AcknowledgementToolTip>;

const Template: ComponentStory<typeof AcknowledgementToolTip> = (args) => (
  <AcknowledgementToolTip {...args} />
);

export const NotAllAcknowledged = Template.bind({});
NotAllAcknowledged.args = {
  data: mockAifRecipient,
};

export const AllAcknowledged = Template.bind({});
AllAcknowledged.args = {
  data: mockAifRecipient2,
};
