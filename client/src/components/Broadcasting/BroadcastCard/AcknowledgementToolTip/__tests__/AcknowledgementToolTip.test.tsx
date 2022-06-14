import { render, screen } from "@testing-library/react";
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
  it("matches snapshot", () => {
    const { container } = render(<AcknowledgementToolTip data={mockData} />);
    expect(container).toMatchSnapshot();
  });
});
