import {
  Card,
  Checkbox,
  CheckboxGroup,
  ActionIcon,
  Group,
  Space,
  Stack,
  Text,
  Textarea,
  TextInput,
  Button,
  ScrollArea,
  Divider,
} from "@mantine/core";
import dayjs from "dayjs";
import React, { useEffect, useState } from "react";
import { Check, Edit, EditOff, Forms } from "tabler-icons-react";
import { useReporting } from "../../../helpers/useReportingClient";
import _ from "lodash";
import { IncidentReport } from "../../../proto/operations_ecosys_pb";
import { useForm } from "@mantine/form";
import { IoClose } from "react-icons/io5";

export default function ReportContainer() {
  const { selectedReport } = useReporting();

  const [editOn, setEditOn] = useState(false);

  const handleEdit = () => {
    console.log("handleEdit called");
    setEditOn(true);
  };

  const handleClose = () => {
    console.log("handleClose called");
    setEditOn(false);
    // TODO: Call updateIncidentReport
  };

  const handleApprove = () => {
    // TODO: Call approveIncidentReport
  };

  useEffect(() => {
    selectedReport &&
      console.log(selectedReport?.incidentReportContent?.hasStolenItem);
    form.setValues({
      title: selectedReport?.incidentReportContent?.title || "",
      address: selectedReport?.incidentReportContent?.address || "",
      time: selectedReport?.incidentReportContent?.incidentTime || "",
      description: selectedReport?.incidentReportContent?.description || "",
      isPoliceNotified:
        selectedReport?.incidentReportContent?.isPoliceNotified || false,
      hasStolenItem:
        selectedReport?.incidentReportContent?.hasStolenItem || false,
    });
  }, [selectedReport]);

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

  return (
    <Card
      sx={{
        flex: "1",
        display: selectedReport?.incidentReportId ? "default" : "none",
        height: "calc(100vh - 132px)",
      }}
    >
      {editOn ? (
        <form
          onSubmit={form.onSubmit((values) => console.log(values))}
          style={{
            height: "100%",
            justifyContent: "space-between",
            display: "flex",
            flexDirection: "column",
          }}
        >
          <Stack>
            <Group position="right">
              <ActionIcon onClick={handleClose}>
                <IoClose />
              </ActionIcon>
            </Group>
            <Stack spacing="lg">
              <Group>
                <Text
                  sx={{
                    width: "125px",
                  }}
                >
                  Title:
                </Text>
                <TextInput
                  sx={{ flex: "1" }}
                  {...form.getInputProps("title")}
                />
              </Group>
              <Group>
                <Text
                  sx={{
                    width: "125px",
                  }}
                >
                  Address:
                </Text>
                <TextInput
                  sx={{ flex: "1" }}
                  {...form.getInputProps("address")}
                />
              </Group>
              <Group>
                <Text
                  sx={{
                    width: "125px",
                  }}
                >
                  Time of incident:
                </Text>
                <TextInput sx={{ flex: "1" }} {...form.getInputProps("time")} />
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
                  {...form.getInputProps("hasStolenItem", { type: "checkbox" })}
                />
              </Group>
            </Stack>
          </Stack>
          <Button color="green" type="submit" fullWidth>
            Save Changes
          </Button>
        </form>
      ) : (
        <Stack spacing={0}>
          <Group position="apart">
            <Text size="lg" weight={500}>
              {selectedReport?.incidentReportContent?.title}
            </Text>
            <ActionIcon onClick={handleEdit}>
              <Edit />
            </ActionIcon>
          </Group>
          <Divider my="sm" />
          <ScrollArea
            sx={{
              height: "480px",
            }}
          >
            <Stack spacing={0} sx={{ fontStyle: "italic" }}>
              <Text size="xs">{`Name: ${selectedReport?.creator?.name}`}</Text>
              <Text size="xs">
                {`Reported on: ${dayjs(
                  selectedReport?.creationDate,
                  "YYYY-MM-DD HH:mm:ss"
                ).format("D/M/YY [at] HH:mm")}`}{" "}
              </Text>
              <Text size="xs">{`Last updated: ${dayjs(
                selectedReport?.lastModifiedDate,
                "YYYY-MM-DD HH:mm:ss"
              ).format("D/M/YY [at] HH:mm")}`}</Text>
              <Text size="xs">{`Address: ${selectedReport?.incidentReportContent?.address}`}</Text>
            </Stack>
            <Space h="md" />
            <Text size="xs">
              {selectedReport?.incidentReportContent?.description}
            </Text>
            <Space h="md" />
            <Stack spacing={0}>
              <Group>
                <Text size="xs" sx={{ width: "120px" }}>
                  Was police notified?
                </Text>
                <Checkbox
                  size="xs"
                  disabled={!editOn}
                  checked={
                    selectedReport?.incidentReportContent?.isPoliceNotified
                  }
                />
              </Group>
              <Group>
                <Text size="xs" sx={{ width: "120px" }}>
                  Was anything stolen?
                </Text>
                <Checkbox
                  size="xs"
                  disabled={!editOn}
                  checked={selectedReport?.incidentReportContent?.hasStolenItem}
                />
              </Group>
            </Stack>
          </ScrollArea>
          <Button
            color="green"
            variant="filled"
            mt="xl"
            onClick={handleApprove}
          >
            Approve
          </Button>
        </Stack>
      )}
    </Card>
  );
}
