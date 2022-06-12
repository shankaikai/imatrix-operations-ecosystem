export {};

describe("broadcasting feature", () => {
  it("should load the page on menu nav", () => {
    cy.visit("/");
    cy.get(".mantine-1myn8qs > :nth-child(1) > :nth-child(2)").click();
  });

  it("should redirect to the right url", () => {
    cy.url().should("include", "broadcasting");
  });

  describe("search bar", () => {
    // Clear search bar after each test
    afterEach(() => {
      cy.get('[data-testid="broadcastSearch"]').clear();
    });

    it("broadcast list should update when search bar is typed", () => {
      const startingState = cy.get('[data-testid="broadcastList"]').children();

      cy.get('[data-testid="broadcastSearch"]')
        .type("Test")
        .should("have.value", "Test");

      const endingState = cy.get('[data-testid="broadcastList"]').children();

      expect(startingState).not.equal(endingState);
    });

    it("broadcast list should have length of 0 when search bar is changed to ---", () => {
      cy.get('[data-testid="broadcastSearch"]')
        .type("---")
        .should("have.value", "---");

      // 1 is for the load more button
      cy.get('[data-testid="broadcastList"]')
        .children()
        .should("have.length", 1);
    });
  });

  describe("sort by", () => {
    it("should reverse the ordering when oldest first option is clicked", () => {
      const startingState = cy.get('[data-testid="broadcastList"]').children();

      cy.get('[data-testid="sortBy"]').click();
      cy.contains("Oldest First").click();

      const endingState = cy.get('[data-testid="broadcastList"]').children();

      expect(endingState).not.eq(startingState);
    });
  });

  describe("new broadcast", () => {
    it("should open modal when new button is clicked", () => {
      cy.get('[data-testid="newBroadcastButton"]').click();
      cy.get(".mantine-Modal-overlay");
    });

    it("should close modal when clicking outside the model", () => {
      cy.get(".mantine-ActionIcon-hover > svg").click();
      cy.get(".mantine-Modal-overlay").should("not.exist");
      cy.wait(1);
    });

    it("should change Recipient inputs", () => {
      cy.get('[data-testid="newBroadcastButton"]').click();
      cy.get('[data-testid="1Badge"]').click();
      cy.get('[data-testid="2Badge"]').click();
    });

    it("should change Urgency inputs", () => {
      cy.get('[data-testid="UrgentBadge"]');
    });

    it("should change messageInput", () => {
      cy.get('[data-testid="newBroadcastMessageInput"]')
        .type(
          "This is a test broadcast. All personnel please report back to HQ immediately."
        )
        .should(
          "have.value",
          "This is a test broadcast. All personnel please report back to HQ immediately."
        );
    });

    it("should submit broadcast", () => {
      cy.get('[data-testid="submitBroadcast"]').click();
    });

    it("should contain the new broadcast on refresh", () => {
      cy.visit("/broadcasting");
      cy.contains(
        "This is a test broadcast. All personnel please report back to HQ immediately."
      );
    });
  });
});
