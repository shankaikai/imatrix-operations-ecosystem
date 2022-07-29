import {
  createContext,
  Dispatch,
  SetStateAction,
  useContext,
  useState,
} from "react";

interface UserInterface {
  name: string;
  email: string;
  image: string;
  userId: number;
  userType: User.UserType;
  setUser?(user: User): void;
}

const UserContext = createContext<UserInterface>({
  name: "",
  email: "",
  image: "",
  userId: -1,
  userType: User.UserType.ISPECIALIST,
});

interface UserProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function UserProvider({ children }: UserProviderProps) {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [image, setImage] = useState("");
  const [userId, setUserId] = useState(-1);
  const [userType, setUserType] = useState<User.UserType>(
    User.UserType.ISPECIALIST
  );
  const setUser = (user: User) => {
    setUserId(user.getUserId());
    setName(user.getName());
    setEmail(user.getEmail());
    setImage(user.getUserSecurityImg());
    setUserType(user.getUserType());
  };
  return (
    <UserContext.Provider
      value={{
        name,
        userId,
        email,
        image,
        userType,
        setUser,
      }}
    >
      {children}
    </UserContext.Provider>
  );
}

export function useUserProvider() {
  return useContext(UserContext);
}
import crypto from "crypto";
import dayjs from "dayjs";
import {
  Client,
  LoginRequest,
  User,
  UserToken,
} from "../proto/operations_ecosys_pb";
import { ENVOY_ADDRESS } from "../utils/constant";
import { AdminServicesClient } from "../proto/Operations_ecosysServiceClientPb";

export const isLoggedIn = () => {
  const jsonString = localStorage.getItem("jwt");
  if (!jsonString) {
    return false;
  }
  const { expiry }: { expiry: string } = JSON.parse(jsonString);
  if (dayjs(expiry).isBefore(dayjs())) {
    return false;
  }

  return true;
};

export const storeJWT = (jwt: string) => {
  localStorage.setItem("jwt", jwt);
};

export const signIn = async (
  username: string,
  password: string
): Promise<User | false> => {
  //TODO: retrieve random string from BE
  const client = getUserClient();
  const user = new User();
  user.setEmail(username);
  const res = await client.getSecurityString(user, {});
  let securityString = res.getSecurityString();

  // append randomString + password
  const hash = crypto
    .createHash("sha256")
    .update(securityString + password)
    .digest("hex");
  //TOOD: send back to backend for validation
  const request = new LoginRequest();
  request.setHashedPassword(hash);
  request.setUserEmail(username);

  let userToken = (await client.authenticateUser(request, {})).getUsertoken();
  if (userToken) {
    const token = userToken.getToken();
    const expiry = userToken.getExpiryDatetime();
    userToken.getUser();
    storeJWT(JSON.stringify({ jwt: token, expiry: expiry }));
    return userToken.getUser() as User;
  } else {
    return false;
  }
};

export const signOut = () => {
  localStorage.removeItem("jwt");
};

export const getUserClient = () => {
  return new AdminServicesClient(ENVOY_ADDRESS, null, {});
};
