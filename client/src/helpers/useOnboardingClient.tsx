import dayjs from "dayjs";
import _ from "lodash";
import {
  createContext,
  Dispatch,
  useContext,
  useEffect,
  useState,
} from "react";
import { RosterServicesClient } from "../proto/Operations_ecosysServiceClientPb";
import {
  AvailabilityQuery,
  BulkRosters,
  EmployeeEvaluation,
  EmployeeEvaluationResponse,
  Filter,
  Roster,
  RosterAssignement,
  RosterFilter,
  RosterQuery,
  RosterResponse,
  User,
} from "../proto/operations_ecosys_pb";
import { ENVOY_ADDRESS } from "../utils/constant";
import getOverallRosterStatus from "./getOverallRosterStatus";
import getRosterDates from "./getRosterDates";
import {
  showErrorNotification,
  showRosterAddSuccessNotification,
  showRosterUpdateSuccessNotification,
} from "./notifications";

enum UserType {
  ISPECIALIST,
  SECURITY_GUARD,
  CONTROLLER,
  MANAGER,
}
interface OnboardingForm {
  user_type: UserType;
  name: string;
  email: string;
  phone_number: string;
  telegram_handle: string;
  user_security_img: string;
  is_part_timer: boolean;
  security_string: string;
  hashed_password: string;
}

export function submitOnboardingForm(onboadingForm: OnboardingForm) {
  const client = getOnboardingClient();

  //todo
}

export function getOnboardingClient(): RosterServicesClient {
  return new RosterServicesClient(ENVOY_ADDRESS, null, {});
}
