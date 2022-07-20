import { Empty } from "google-protobuf/google/protobuf/empty_pb";
import {
  createContext,
  Dispatch,
  useContext,
  useEffect,
  useState,
} from "react";
import { CameraIotServicesClient } from "../proto/Operations_ecosysServiceClientPb";
import { CameraIot, CameraIotResponse } from "../proto/operations_ecosys_pb";

interface CameraIotInterface {
  search: string;
  setSearch?: Dispatch<string>;
  cameras: CameraIot.AsObject[];
  setCameras?: Dispatch<CameraIot.AsObject[]>;
}

const CameraIotContext = createContext<CameraIotInterface>({
  search: "",
  cameras: [],
});

interface CameraIotProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function CameraIotProvider({ children }: CameraIotProviderProps) {
  const [search, setSearch] = useState<string>("");
  const [cameras, setCameras] = useState<CameraIot.AsObject[]>([]);

  useEffect(() => {
    getCameraFeeds(setCameras);
  }, []);

  return (
    <CameraIotContext.Provider
      value={{ search, setSearch, cameras, setCameras }}
    >
      {children}
    </CameraIotContext.Provider>
  );
}

export function getCameraIotClient(): CameraIotServicesClient {
  return new CameraIotServicesClient("http://localhost:8080", null);
}

export function useCameraIot() {
  return useContext(CameraIotContext);
}

export function getCameraFeeds(setCameras: Dispatch<CameraIot.AsObject[]>) {
  console.log("getCameraFeeds called");

  const client = getCameraIotClient();

  const stream = client.getIotState(new Empty());

  stream.on("data", (response: CameraIotResponse) => {
    console.log(response.getCameraIot()?.toObject());
    //@ts-ignore
    setCameras((oldState) => [
      ...oldState,
      response.getCameraIot()?.toObject(),
    ]);
  });
}
