import { BroadcastServicesClient } from "../proto/operations_ecosys_grpc_web_pb";

export default function useBroadcastClient(): BroadcastServicesClient {
  // TODO: add the envoy address into .env
  return new BroadcastServicesClient("http://localhost:8080", null, {});
}
