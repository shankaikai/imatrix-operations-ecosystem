import {
  ActionIcon,
  Card,
  Checkbox,
  CheckboxGroup,
  Group,
  Space,
  Stack,
  Text,
} from "@mantine/core";
import React from "react";
import { IoSave } from "react-icons/io5";

export default function ReportContainer() {
  const handleSave = () => {
    console.log("handleSave called");
  };

  return (
    <Card
      sx={{
        width: "50%",
      }}
    >
      <Stack>
        <Group position="apart">
          <Text size="lg" weight={500}>
            Camera Faulty at NOK Site
          </Text>
          <ActionIcon onClick={handleSave}>
            <IoSave />
          </ActionIcon>
        </Group>
        <Space />
        <Stack spacing={0}>
          <Text size="xs">{`Name:`}</Text>
          <Text size="xs">{`Reported on:`}</Text>
          <Text size="xs">{`Last updated:`}</Text>
          <Text size="xs">{`Address:`}</Text>
        </Stack>
        <Space />
        <Text size="xs">
          Aliqua id fugiat nostrud irure ex duis ea quis id quis ad et. Sunt qui
          esse pariatur duis deserunt mollit dolore cillum minim tempor enim.
          Elit aute irure tempor cupidatat incididunt sint deserunt ut voluptate
          aute id deserunt nisi. Pasd lasdnpwasd pwoel 12pm asknxcloa oasjdpiwh
          pasijndasd noianc pojposjadnwpo asopjd mkxcklwpojd poa aojapsodj
          apojpoawjdpias oapwjd asojd asjd p osjf sdif soefj poc kashdoiawhdp
          flm spd aoj w0e d asd axc es f rerwet oto so spdfk sodvxclv nlxckv
          psoefwe sadicspodf we0u apsdo jsdpfo jpdouw e0-s jdsdosfj woe0u=-ue0=
          ew jwopdajpodjpo fop spof 0e if0sd-fi -0sif 0sier 0qir 39-r uq9e u
          w0ri -0e -0sdfj aop 0ur9 selkfsjdpfdirg r0soeip suep s9eoe ir 0sp9f
          pse 9f. Aliqua id fugiat nostrud irure ex duis ea quis id quis ad et.
          Sunt qui esse pariatur duis deserunt mollit dolore cillum minim tempor
          enim. Elit aute irure tempor cupidatat. Aliqua id fugiat nostrud irure
          ex duis ea quis id quis ad et. Sunt qui esse pariatur duis deserunt
          mollit dolore cillum minim tempor enim. Elit aute irure tempor
          cupidatat
        </Text>
        <Space />
        <CheckboxGroup>
          <Checkbox label="Test" />
          <Checkbox />
        </CheckboxGroup>
      </Stack>
    </Card>
  );
}
