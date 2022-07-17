import { createContext, useContext } from "react";
import { CameraIotServicesClient } from "../proto/Operations_ecosysServiceClientPb";

interface CameraIotInterface {}

const CameraIotContext = createContext<CameraIotInterface>({});

interface CameraIotProviderProps {
  children: JSX.Element | JSX.Element[];
}

export function CameraIotProvider({ children }: CameraIotProviderProps) {
  return (
    <CameraIotContext.Provider value={{}}>{children}</CameraIotContext.Provider>
  );
}

export function getReportingClient(): CameraIotServicesClient {
  return new CameraIotServicesClient("http://localhost:8080", null);
}

export function useCameraIot() {
  return useContext(CameraIotContext);
}
