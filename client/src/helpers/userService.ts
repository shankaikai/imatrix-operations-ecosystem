import crypto from "crypto";
import dayjs from "dayjs";

export const isLoggedIn = () => {
  const jsonString = localStorage.getItem("jwt");
  if (!jsonString) {
    return false;
  }
  const { expiry }: { expiry: string } =
    JSON.parse(jsonString);
  if (dayjs(expiry).isBefore(dayjs())) {
    return false;
  }
  return true;
};

export const storeJWT = (jwt: string) => {
  localStorage.setItem("jwt", jwt);
};

export const signIn = (username: string, password: string): boolean => {
  //TODO: retrieve random string from BE
  const randomString = "abcde";
  // append randomString + password
  const hash = crypto.createHash("sha256").update(randomString + password);

  //TOOD: send back to backend for validation
  const jwt = "JWT";
  const expiry = dayjs().add(1,'day').toString();
  if (jwt) {
    storeJWT(JSON.stringify({ jwt: jwt, expiry: expiry }));
    return true;
  } else {
    return false;
  }
};

export const signOut = () => {
  localStorage.removeItem("jwt");
}