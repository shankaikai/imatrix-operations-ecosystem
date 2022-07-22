import {
  ActionIcon,
  Text,
  Button,
  Checkbox,
  Group,
  Input,
  Modal,
  Stack,
  Textarea,
  TextInput,
} from "@mantine/core";
import { useForm } from "@mantine/hooks";
import { FunctionComponent, useEffect } from "react";
import { IoClose } from "react-icons/io5";
import {
  submitUpdateReport,
  UpdateReport,
  useReporting,
} from "../../../helpers/useReportingClient";
import { useUserProvider } from "../../../helpers/useUserProvider";
import { IncidentReport } from "../../../proto/operations_ecosys_pb";

interface ReportsModalProps {}

const ReportsModal: FunctionComponent<ReportsModalProps> = () => {
  const { modalOpen, setModalOpen, createNewReport, setReports } =
    useReporting();

  const { userId } = useUserProvider();
  const form = useForm({
    initialValues: {
      title: "",
      address: "",
      time: "",
      description: "",
      isPoliceNotified: false,
      hasStolenItem: false,
    },
  });

  const closeModal = () => {
    setModalOpen && setModalOpen(false);
  };

  const onSubmit = async (values: UpdateReport) => {
    setReports &&
      createNewReport &&
      (await createNewReport(values, setReports, userId));
    closeModal();
  };
  return (
    <Modal title="New Report" opened={modalOpen} onClose={closeModal} size="xl">
      <form
        onSubmit={form.onSubmit((values) => onSubmit(values))}
        style={{
          height: "100%",
          justifyContent: "space-between",
          display: "flex",
          flexDirection: "column",
          width: "100%",
        }}
      >
        <Stack spacing="xl">
          <Group>
            <Text
              sx={{
                width: "125px",
              }}
            >
              Title:
            </Text>
            <TextInput sx={{ flex: "1" }} {...form.getInputProps("title")} />
          </Group>
          <Group>
            <Text
              sx={{
                width: "125px",
              }}
            >
              Address:
            </Text>
            <TextInput sx={{ flex: "1" }} {...form.getInputProps("address")} />
          </Group>
          <Group>
            <Text
              sx={{
                width: "125px",
              }}
            >
              Time of incident:
            </Text>
            <Input
              type="datetime-local"
              sx={{ flex: "1" }}
              {...form.getInputProps("time")}
            />
          </Group>
          <Group
            sx={{
              alignItems: "flex-start",
            }}
          >
            <Text
              sx={{
                width: "125px",
              }}
            >
              Details:
            </Text>
            <Textarea
              sx={{ flex: "1" }}
              autosize
              maxRows={5}
              {...form.getInputProps("description")}
            />
          </Group>
          <Group>
            <Text
              sx={{
                width: "125px",
              }}
            >
              Police Notified:
            </Text>
            <Checkbox
              {...form.getInputProps("isPoliceNotified", {
                type: "checkbox",
              })}
            />
          </Group>
          <Group>
            <Text
              sx={{
                width: "125px",
              }}
            >
              Stolen Items:
            </Text>
            <Checkbox
              {...form.getInputProps("hasStolenItem", {
                type: "checkbox",
              })}
            />
          </Group>
        </Stack>
        <Button color="green" type="submit" mt="xl" fullWidth>
          Save Changes
        </Button>
      </form>
    </Modal>
  );
};

export default ReportsModal;
