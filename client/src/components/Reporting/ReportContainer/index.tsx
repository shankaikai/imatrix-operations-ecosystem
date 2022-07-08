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
  Loader,
  Input,
} from "@mantine/core";
import dayjs from "dayjs";
import React, { useEffect, useState } from "react";
import { Check, Edit, EditOff, Forms } from "tabler-icons-react";
import {
  approveReport,
  submitUpdateReport,
  useReporting,
} from "../../../helpers/useReportingClient";
import _ from "lodash";
import { IncidentReport } from "../../../proto/operations_ecosys_pb";
import { useForm } from "@mantine/form";
import { IoClose } from "react-icons/io5";

export default function ReportContainer() {
  const { selectedReport, setSelectedReport, setReports, reports } =
    useReporting();

  const [editOn, setEditOn] = useState(false);

  const handleEdit = () => {
    console.log("handleEdit called");
    setEditOn(true);
  };

  const handleClose = () => {
    console.log("handleClose called");
    setEditOn(false);
  };

  const handleApprove = async () => {
    const id =
      selectedReport &&
      setReports &&
      (await approveReport(selectedReport.incidentReportId, setReports));
    setSelectedReport &&
      setSelectedReport(
        reports.find(
          (report) => report.incidentReportId === id
        ) as IncidentReport.AsObject
      );
    forceUpdate();
  };

  const [, updateState] = React.useState();
  //@ts-ignore
  const forceUpdate = React.useCallback(() => updateState({}), []);

  useEffect(() => {
    selectedReport &&
      form.setValues({
        title: selectedReport?.incidentReportContent?.title || "",
        address: selectedReport?.incidentReportContent?.address || "",
        time:
          dayjs(selectedReport?.incidentReportContent?.incidentTime).format(
            "YYYY-MM-DD[T]HH:mm"
          ) || "",
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

  return selectedReport ? (
    <Card
      sx={{
        flex: "1",
        display: selectedReport?.incidentReportId ? "flex" : "none",
        height: "calc(100vh - 132px)",
        width:'100%'
      }}
    >
      {editOn ? (
        <form
          onSubmit={form.onSubmit((values) =>
            submitUpdateReport(
              values,
              selectedReport?.incidentReportId as number,
              setReports as React.Dispatch<IncidentReport.AsObject[]>
            )
          )}
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
          </Stack>
          <Button color="green" type="submit" fullWidth>
            Save Changes
          </Button>
        </form>
      ) : (
        <Stack spacing={0} sx={{width:'100%'}}>
          <Group position="apart">
            <Text size="lg" weight={500}>
              {selectedReport?.incidentReportContent?.title}
            </Text>
            {!selectedReport.isApproved && (
              <ActionIcon onClick={handleEdit}>
                <Edit />
              </ActionIcon>
            )}
          </Group>
          <Divider my="sm" />
          <ScrollArea
            sx={{
              flex:'1',
            }}
          >
            <Stack spacing={0} sx={{ fontStyle: "italic" }}>
              <Text size="xs">{`Name: ${selectedReport?.creator?.name}`}</Text>
              <Text size="xs">
                {`Reported on: ${dayjs(
                  selectedReport?.incidentReportContent?.incidentTime,
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
          {!selectedReport.isApproved && (
            <Button
              color="green"
              variant="filled"
              mt="xl"
              onClick={handleApprove}
            >
              Approve
            </Button>
          )}
        </Stack>
      )}
    </Card>
  ) : (
    <div />
  );
}
