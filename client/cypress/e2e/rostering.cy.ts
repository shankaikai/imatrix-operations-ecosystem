describe("rostering feature", () => {
  it("should load the page on menu nav", () => {
    cy.visit("/");
    cy.get(".mantine-1myn8qs > :nth-child(1) > :nth-child(3)").click();
  });

  it("should go to next week on arrow click", () => {
    cy.get(":nth-child(8) > .icon").click();
  });

  it("should change selected date", () => {
    cy.wait(2000);
    cy.get(".mantine-19jxmdp > :nth-child(2)").click();
    cy.wait(2000);
    cy.get(".mantine-19jxmdp > :nth-child(4)").click();
  });

  it("should replace aifs basket on drag and drop", () => {
    cy.get(":nth-child(6) > .mantine-Popover-root").trigger("dragstart");
    cy.get(
      ':nth-child(1) > .mantine-lfk3cq > [style="border: 2px dashed rgb(173, 181, 189); height: auto; padding: 10px; border-radius: 12px; display: flex; align-items: center; justify-content: center;"]'
    ).trigger("drop");

    cy.get(":nth-child(4) > .mantine-Popover-root").trigger("dragstart");
    cy.get(
      ':nth-child(2) > .mantine-lfk3cq > [style="border: 2px dashed rgb(173, 181, 189); height: auto; padding: 10px; border-radius: 12px; display: flex; align-items: center; justify-content: center;"]'
    ).trigger("drop");
  });

  it("should publish on publish click", () => {
    try {
      cy.get(".mantine-2phkfh").click();
    } catch {}
    cy.get(".mantine-2phkfh").should("be.disabled");
  });
});
