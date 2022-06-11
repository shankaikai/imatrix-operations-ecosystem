import { Dispatch } from "react";
import { render } from "../../../../utils/test-utils";
import BroadcastFilter from "../BroadcastFilter";

const mockDispatch = jest.fn();

describe("BroadcastFilter", () => {
  it("matches snapshot", () => {
    const { container } = render(
      <BroadcastFilter setModalOpen={mockDispatch} />
    );

    expect(container).toMatchSnapshot();
  });
});
