import {
  createContext,
  Dispatch,
  SetStateAction,
  useContext,
  useState,
} from "react";

interface UserInterface {
  name: string;
  setName?: Dispatch<SetStateAction<string>>;
  email: string;
  setEmail?: Dispatch<SetStateAction<string>>;
  image: string;
  setImage?: Dispatch<SetStateAction<string>>;
}

const UserContext = createContext<UserInterface>({
  name: "",
  email: "",
  image: "",
});

interface UserProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function UserProvider({ children }: UserProviderProps) {
  const [name, setName] = useState("PH Chang");
  const [email, setEmail] = useState("ph@imatrix.sg");
  const [image, setImage] = useState(
    "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
  );

  return (
    <UserContext.Provider
      value={{
        name,
        setName,
        email,
        setEmail,
        image,
        setImage,
      }}
    >
      {children}
    </UserContext.Provider>
  );
}

export function useUserProvider() {
  return useContext(UserContext);
}
