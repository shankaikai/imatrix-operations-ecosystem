import { render } from "@testing-library/react";
import {
  AIFSRecipient,
  RecipientDetails,
} from "../../../../helpers/recipientsFormatter";
import BroadcastCard from "../BroadcastCard";

const mockContent = "some content";
const mockDate = "some date";
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
const mockAifs = [mockData];

describe("BroadcastCard", () => {
  it("matches snapshot", () => {
    const { container } = render(
      <BroadcastCard content={mockContent} date={mockDate} aifs={mockAifs} />
    );
    expect(container).toMatchSnapshot();
  });
});
