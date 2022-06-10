import { render } from "@testing-library/react";
import AcknowledgementToolTip from "..";
import {
  AIFSRecipient,
  RecipientDetails,
} from "../../../../../helpers/recipientsFormatter";

const mockRecipientDetails: RecipientDetails[] = [
  {
    id: 1,
    phone: "some phone number",
    acknowledged: true,
    img: "some url",
    name: "some name",
  },
];

const mockData: AIFSRecipient = {
  id: "some id",
  location: "some location",
  allAcknowledged: true,
  users: mockRecipientDetails,
};

describe("AcknowledgementToolTip", () => {
  it("renders all components correctly", () => {
    render(<AcknowledgementToolTip data={mockData} />);

    // expect()
  });
});
