import { Empty } from "google-protobuf/google/protobuf/empty_pb";
import { createContext, useContext, useEffect, useState } from "react";
import { CameraIotServicesClient } from "../proto/Operations_ecosysServiceClientPb";
import { CameraIotResponse } from "../proto/operations_ecosys_pb";

interface CameraIotInterface {}

const CameraIotContext = createContext<CameraIotInterface>({});

interface CameraIotProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function CameraIotProvider({ children }: CameraIotProviderProps) {
  const [search, setSearch] = useState<string>("");

  useEffect(() => {
    getCameraFeeds();
  });

  return (
    <CameraIotContext.Provider value={{ search, setSearch }}>
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

export function getCameraFeeds() {
  console.log("getCameraFeeds called");

  const client = getCameraIotClient();

  const stream = client.getIotState(new Empty());

  stream.on("data", (response: CameraIotResponse) => {
    console.log(response);
  });
}
