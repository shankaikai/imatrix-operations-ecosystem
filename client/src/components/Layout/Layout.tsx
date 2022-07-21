import { AppShell, ScrollArea, useMantineColorScheme } from "@mantine/core";
import MainHeader from "./MainHeader";
import Sidebar from "./Sidebar/Sidebar";

type LayoutProps = {
  children: React.ReactNode;
};

export default function Layout({ children }: LayoutProps) {
  return (
    <AppShell
      padding="md"
      navbar={<Sidebar />}
      header={<MainHeader />}
      styles={(theme) => ({
        main: {
          backgroundColor:
            theme.colorScheme === "dark"
              ? theme.colors.dark[8]
              : theme.colors.gray[0],
          padding: theme.spacing.xs,
        },
      })}
      fixed
    >
      <div
        style={{
          height: "calc 100vh - 60px",
          marginLeft: 280,
          marginTop: 60,
          position: "relative",
        }}
      >
        {children}
      </div>
    </AppShell>
  );
}
