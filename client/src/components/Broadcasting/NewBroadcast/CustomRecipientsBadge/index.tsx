import { Badge } from "@mantine/core";
import React, { Dispatch } from "react";

interface CustomRecipientsBadgeProps {
  label: string;
  value: string;
  setValue: Dispatch<any>;
  active?: boolean;
}

export default function CustomRecipientsBadge({
  label,
  value,
  setValue,
  active,
}: CustomRecipientsBadgeProps) {
  const handleClick = () => {
    if (value === "all") {
      if (active) {
        setValue(["all"]);
      } else {
        setValue([]);
      }
    }

    if (active) {
      setValue((prevValue) => prevValue.filter((e) => e !== value));
    } else {
      setValue((prevValue) => [...prevValue, value]);
    }
  };

  return (
    <Badge
      data-testid={`${value}Badge`}
      id={value}
      sx={{
        cursor: "pointer",
      }}
      onClick={() => handleClick()}
      variant={active ? "filled" : "light"}
    >
      {label}
    </Badge>
  );
}
