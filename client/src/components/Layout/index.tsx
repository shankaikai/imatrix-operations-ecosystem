import { createStyles, Container } from "@mantine/core";
import { Sidebar } from "./Sidebar";

const useStyles = createStyles((theme) => ({
  layout: {
    display: "flex",
    flexDirection: "row",
    height: "100vh",
    width: "100vw",
  },
  content: {
    background: theme.colors.gray[1],
    width: "100%",
  },
}));

type LayoutProps = {
  children: React.ReactNode;
};

export default function Layout({ children }: LayoutProps) {
  const { classes } = useStyles();

  return (
    <div className={classes.layout}>
      <Sidebar />
      <div className={classes.content}>
        <main>{children}</main>
      </div>
    </div>
  );
}
