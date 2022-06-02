import { AIFSBroadcastRecipient } from "../proto/operations_ecosys_pb";

export interface RecipientDetails {
  id?: number;
  phone?: string;
  acknowledged: boolean;
  img?: string;
  name?: string;
}

export interface AIFSRecipient {
  id: string;
  location: string;
  allAcknowledged: boolean;
  users: RecipientDetails[];
}
export default function recipientFormatter(
  recipients: AIFSBroadcastRecipient[]
) {
  const out: AIFSRecipient[] = [];

  for (var recipient of recipients) {
    const workers = recipient.getRecipientList();
    var users = [];
    var allAcknowledged = true;

    for (var worker of workers) {
      const user: RecipientDetails = {
        id: worker.getRecipient()?.getUserId(),
        phone: worker.getRecipient()?.getPhoneNumber(),
        acknowledged: worker.getAcknowledged(),
        img: worker.getRecipient()?.getUserSecurityImg(),
        name: worker.getRecipient()?.getName(),
      };

      users.push(user);
      allAcknowledged = allAcknowledged && worker.getAcknowledged();
    }

    const aifs: AIFSRecipient = {
      id: "AIFS " + recipient.getAifsId().toString(),
      location: "Location",
      users,
      allAcknowledged,
    };

    out.push(aifs);
  }

  return out;
}
