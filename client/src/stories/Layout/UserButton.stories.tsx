import { ComponentMeta, ComponentStory } from "@storybook/react";
import { UserButton } from "../../components/Layout/UserButton";

export default {
  title: "Shared/UserButton",
  component: UserButton,
} as ComponentMeta<typeof UserButton>;

const Template: ComponentStory<typeof UserButton> = (args) => (
  <UserButton {...args} />
);

export const Default = Template.bind({});
Default.args = {
  image: "https://picsum.photos/202",
  name: "Some Name",
  email: "someemail@gmail.com",
};
