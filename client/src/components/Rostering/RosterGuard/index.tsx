import { Avatar } from "@mantine/core";
import React from "react";

interface RosterGuardProps {
  id: number;
  img: string;
  name: string;
  phone: string;
}

export default function RosterGuard({
  id,
  img,
  name,
  phone,
}: RosterGuardProps) {
  return (
    <Avatar
      size={80}
      src="https://www.khaosodenglish.com/wp-content/uploads/2020/02/guard-copy.jpg"
      radius={100}
    />
  );
}
