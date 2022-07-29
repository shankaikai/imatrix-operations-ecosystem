import { ComponentStory, ComponentMeta } from "@storybook/react";
import ReportSmallCard from "../../components/Reporting/ReportSmallCard";

export default {
  title: "Reporting/ReportSmallCard",
  component: ReportSmallCard,
} as ComponentMeta<typeof ReportSmallCard>;

const Template: ComponentStory<typeof ReportSmallCard> = (args) => (
  <ReportSmallCard {...args} />
);

export const Default = Template.bind({});
Default.args = {
  id: 1,
  title: "Some Title",
  sender: "Sender",
  senderImg: "https://picsum.photos/200",
  aifsId: 1,
  creationDate: "2022-10-22 11:10:40",
  updateDate: "2022-10-22 11:10:20",
  selected: false,
};

export const Selected = Template.bind({});
Selected.args = {
  id: 1,
  title: "Some Title",
  sender: "Sender",
  senderImg: "https://picsum.photos/200",
  aifsId: 1,
  creationDate: "2022-10-22 11:10:40",
  updateDate: "2022-10-22 11:10:20",
  selected: true,
};
