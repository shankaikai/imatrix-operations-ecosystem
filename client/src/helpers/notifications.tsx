import { showNotification } from "@mantine/notifications";
import { ExclamationMark } from "tabler-icons-react";
import { IoCheckmark } from "react-icons/io5";

export function showBroadcastSuccessNotification() {
  showNotification({
    title: "Success",
    message: "Broadcast has been published successfully",
    color: "green",
    icon: <IoCheckmark />,
  });
}

export function showErrorNotification() {
  showNotification({
    title: "Error",
    message: "Something went wrong",
    color: "red",
    icon: <ExclamationMark />,
  });
}

export function showRosterAddSuccessNotification() {
  showNotification({
    title: "Success",
    message: "Roster has been published successfully",
    color: "green",
    icon: <IoCheckmark />,
  });
}

export function showRosterUpdateSuccessNotification() {
  showNotification({
    title: "Success",
    message: "Roster has been updated successfully",
    color: "green",
    icon: <IoCheckmark />,
  });
}

export function showUpdateReportSuccessNotification() {
  showNotification({
    title: "Success",
    message: "Report has been updated successfully",
    color: "green",
    icon: <IoCheckmark />,
  });
}

export function showApproveReportSuccessNotification() {
  showNotification({
    title: "Success",
    message: "Report has been approved successfully",
    color: "green",
    icon: <IoCheckmark />,
  });
}
export function showCreateUserSuccessNotification() {
  showNotification({
    title: "Success",
    message: "User has been created!",
    color: "green",
    icon: <IoCheckmark />,
  });
}
