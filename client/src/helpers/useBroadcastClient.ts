import { BroadcastServicesClient } from "../proto/Operations_ecosysServiceClientPb";

export default function getBroadcastClient(): BroadcastServicesClient {
  // TODO: add the envoy address into .env
  return new BroadcastServicesClient("http://localhost:8080", null, {});
}
