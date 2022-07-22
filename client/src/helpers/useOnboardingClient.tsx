import _ from "lodash";
import { AdminServicesClient } from "../proto/Operations_ecosysServiceClientPb";
import { FullUser, User } from "../proto/operations_ecosys_pb";
import { ENVOY_ADDRESS } from "../utils/constant";
import { showCreateUserSuccessNotification } from "./notifications";
import crypto from "crypto";

enum UserType {
  ISPECIALIST,
  SECURITY_GUARD,
  CONTROLLER,
  MANAGER,
}
export interface OnboardingForm {
  user_type: UserType;
  name: string;
  email: string;
  phone_number: string;
  telegram_handle: string;
  user_security_img: string;
  is_part_timer: boolean;
  password: string;
}

export function submitOnboardingForm(onboardingForm: OnboardingForm) {
  const client = getOnboardingClient();

  const user = new User();

  debugger;
  user.setEmail(onboardingForm.email);
  user.setName(onboardingForm.name);
  user.setPhoneNumber(onboardingForm.phone_number);
  user.setTelegramHandle(onboardingForm.telegram_handle);
  user.setUserType(onboardingForm.user_type);
  user.setUserSecurityImg(onboardingForm.user_security_img);
  user.setIsPartTimer(onboardingForm.is_part_timer);

  const request = new FullUser();

  const securityString = (Math.random() * 64).toString(36);
  const hash = crypto
    .createHash("sha256")
    .update(securityString + onboardingForm.password)
    .digest("hex");

  request.setSecurityString(securityString);
  request.setUser(user);
  request.setHashedPassword(hash);

  debugger;
  client
    .addUser(request, {})
    .then((response) => {
      showCreateUserSuccessNotification();
    })
    .catch((e) => console.log(e));
}

export function getOnboardingClient(): AdminServicesClient {
  return new AdminServicesClient(ENVOY_ADDRESS, null, {});
}
