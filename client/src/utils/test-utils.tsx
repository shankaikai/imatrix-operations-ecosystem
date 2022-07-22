import React, { FC, ReactElement } from "react";
import { render, RenderOptions } from "@testing-library/react";
import { ColorSchemeProvider, MantineProvider } from "@mantine/core";

const AllTheProviders: FC<{ children: React.ReactNode }> = ({ children }) => {
  const toggleColorScheme = () => {};

  return (
    <ColorSchemeProvider
      colorScheme="light"
      toggleColorScheme={toggleColorScheme}
    >
      <MantineProvider>{children}</MantineProvider>
    </ColorSchemeProvider>
  );
};

const customRender = (
  ui: ReactElement,
  options?: Omit<RenderOptions, "wrapper">
) => render(ui, { wrapper: AllTheProviders, ...options });

export * from "@testing-library/react";
export { customRender as render };
