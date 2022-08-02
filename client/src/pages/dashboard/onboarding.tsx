import {
  TextInput,
  Checkbox,
  Button,
  Group,
  Box,
  Stack,
  Select,
} from "@mantine/core";
import type { NextPage } from "next";
import Layout from "../../components/Layout/Layout";
import {
  OnboardingForm,
  submitOnboardingForm,
} from "../../helpers/useOnboardingClient";
import { useForm } from "@mantine/form";
import { User } from "../../proto/operations_ecosys_pb";

const Reporting: NextPage = () => {
  const form = useForm<OnboardingForm>({
    initialValues: {
      user_type: User.UserType.CONTROLLER,
      email: "",
      name: "",
      phone_number: "",
      telegram_handle: "",
      user_security_img: "",
      is_part_timer: false,
      password: "",
    },

    validate: {
      email: (value) => (/^\S+@\S+$/.test(value) ? null : "Invalid email"),
    },
  });

  return (
    <Layout>
      <Stack>
        <form
          onSubmit={form.onSubmit((values) => submitOnboardingForm(values))}
        >
          <Select
            required
            label="User Role"
            data={[
              {
                value: User.UserType.CONTROLLER.toString(),
                label: "Controller",
              },
              {
                value: User.UserType.ISPECIALIST.toString(),
                label: "I-Specialist",
              },
              {
                value: User.UserType.SECURITY_GUARD.toString(),
                label: "Security Guard",
              },
              { value: User.UserType.MANAGER.toString(), label: "Manager" },
            ]}
            onChange={form.getInputProps("user_type").onChange}
          />
          <TextInput required label="Name" {...form.getInputProps("name")} />

          <TextInput
            required
            label="Phone number"
            {...form.getInputProps("phone_number")}
          />

          <TextInput
            required
            label="Telegram handle"
            {...form.getInputProps("telegram_handle")}
          />

          <TextInput
            required
            type="password"
            label="Password"
            {...form.getInputProps("password")}
          />
          <TextInput required label="Email" {...form.getInputProps("email")} />

          <Checkbox
            mt="md"
            label="Part Timer"
            {...form.getInputProps("is_part_timer", { type: "checkbox" })}
          />

          <Group position="right" mt="md">
            <Button type="submit">Submit</Button>
          </Group>
        </form>
      </Stack>
    </Layout>
  );
};

export default Reporting;
