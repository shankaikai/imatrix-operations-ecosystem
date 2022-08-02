import type { NextPage } from "next";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { isLoggedIn, useUserProvider } from "../helpers/useUserProvider";
import { User } from "../proto/operations_ecosys_pb";

const Home: NextPage = () => {
  const router = useRouter();
  const { setUser } = useUserProvider();

  useEffect(() => {
    const user = isLoggedIn();
    if (!user) {
      router.push("/login");
    } else {
      setUser && setUser(user);
      router.push("/dashboard/broadcasting");
    }
  }, []);
  return <></>;
};

export default Home;
