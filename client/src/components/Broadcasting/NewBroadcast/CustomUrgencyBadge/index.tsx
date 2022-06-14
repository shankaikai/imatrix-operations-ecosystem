import { Badge } from "@mantine/core";
import React, { Dispatch } from "react";

interface CustomUrgencyBadgeProps {
  label: string;
  setValue: Dispatch<any>;
  active?: boolean;
}

export default function CustomUrgencyBadge({
  label,
  setValue,
  active,
}: CustomUrgencyBadgeProps) {
  const handleClick = () => {
    if (active) {
      setValue("");
    } else {
      setValue(label);
    }
  };

  return (
    <Badge
      data-testid={`${label}Badge`}
      id={label}
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
