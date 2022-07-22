import type { NextPage } from "next";
import { useRouter } from "next/router";
import { useEffect } from "react";
import { isLoggedIn } from "../helpers/useUserProvider";


const Home: NextPage = () => {

  const router = useRouter();

  useEffect(() => {
    if (!isLoggedIn()) {
        router.push('/login');
    } else {
        router.push('/dashboard')
    }
}, []);
return (
    <>
    </>
  );
};

export default Home;
