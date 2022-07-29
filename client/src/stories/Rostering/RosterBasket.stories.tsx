import { ComponentMeta, ComponentStory } from "@storybook/react";
import RosterBasket from "../../components/Rostering/RosterBasket";

export default {
  title: "Rostering/RosterBasket",
  component: RosterBasket,
} as ComponentMeta<typeof RosterBasket>;

const Template: ComponentStory<typeof RosterBasket> = (args) => (
  <RosterBasket {...args} />
);

export const Default = Template.bind({});
Default.args = {};
