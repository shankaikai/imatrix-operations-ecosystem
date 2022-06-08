import type { NextPage } from "next";
import RosterContainer from "../components/Rostering/RosterContainer";
import { RosteringProvider } from "../helpers/useRosteringClient";

// const aifs = ["AIFS 1 (AMKC)", "AIFS 2 (BKP)", "AIFS 3 (PKC)"];

// const guards = [
//   {
//     id: 1,
//     name: "Guard1",
//     img: "https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg",
//     phone: "92818838",
//   },
//   {
//     id: 2,
//     name: "Guard2",
//     img: "https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg",
//     phone: "92818838",
//   },
//   {
//     id: 3,
//     name: "Guard3",
//     img: "https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg",
//     phone: "92818838",
//   },
// ];

const Rostering: NextPage = () => {
  return (
    <RosteringProvider>
      <RosterContainer />
    </RosteringProvider>
  );
};

export default Rostering;
