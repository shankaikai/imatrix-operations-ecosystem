import { createStyles, Container } from "@mantine/core";
import { Sidebar } from "./Sidebar";

const useStyles = createStyles((theme) => ({
  layout: {
    display: "flex",
    flexDirection: "row",
    height: "vh",
    width: "vw",
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
      <Container>
        <main>{children}</main>
      </Container>
    </div>
  );
}
