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
import _ from "lodash";
import { GateState } from "../proto/iot_prototype_pb";
import { ENVOY_ADDRESS } from "../utils/constant";

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
  return new CameraIotServicesClient(ENVOY_ADDRESS, null);
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
    const type = response.getCameraIot()?.getType();
    const cameraIot = response.getCameraIot()?.toObject();

    if (type === CameraIot.MessageType.INITIAL) {
      console.log("inital");
      //@ts-ignore
      setCameras((oldState) => [...oldState, cameraIot]);
    } else if (type === CameraIot.MessageType.CHANGE_CPU_TEMP) {
      console.log("cpu temp");

      //@ts-ignore
      setCameras((oldState: CameraIot.AsObject[]) => {
        let newState = _.cloneDeep(oldState);
        let cam = newState.find(
          (camera) => camera.cameraIotId === cameraIot?.cameraIotId
        );
        if (cam) {
          cam.cpuTemperature = cameraIot?.cpuTemperature;
        }
        return newState;
      });
    } else if (type === CameraIot.MessageType.CHANGE_FIRE_ALARM) {
      console.log("fire alarm");

      //@ts-ignore
      setCameras((oldState: CameraIot.AsObject[]) => {
        let newState = _.cloneDeep(oldState);
        let cam = newState.find(
          (camera) => camera.cameraIotId === cameraIot?.cameraIotId
        );
        if (cam) {
          cam.fireAlarm = cameraIot?.fireAlarm;
        }
        return newState;
      });
    } else if (type === CameraIot.MessageType.CHANGE_GATE) {
      console.log("change gate");

      //@ts-ignore
      setCameras((oldState: CameraIot.AsObject[]) => {
        let newState = _.cloneDeep(oldState);
        let cam = newState.find(
          (camera) => camera.cameraIotId === cameraIot?.cameraIotId
        );
        if (cam) {
          cam.gate = cameraIot?.gate;
        }
        return newState;
      });
    }
  });
}

export function activateGateSwitch(
  id: number,
  gateStatus: GateState.GatePosition
) {
  const client = getCameraIotClient();

  const gateState = new GateState();

  gateStatus === GateState.GatePosition.CLOSED
    ? gateState.setState(GateState.GatePosition.OPEN)
    : gateState.setState(GateState.GatePosition.CLOSED);

  gateState.setId(id);

  client
    .setGateState(gateState, null)
    .then((response) => {
      //TODO: Notification
      console.log(response);
    })
    .catch((error) => {
      console.log(error);
    });
}
