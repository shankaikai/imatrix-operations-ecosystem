import { FunctionComponent, useEffect, useState } from "react";
import { Card, Text, TextInput, Checkbox, Button, Group, Box } from '@mantine/core';
import {  } from '@mantine/core';
import { useForm } from '@mantine/form';
import { isLoggedIn, signIn } from "../helpers/userService";
import { useRouter } from "next/router";

interface LogInProps {
    
}
 
const LogIn: FunctionComponent<LogInProps> = () => {

    const router = useRouter();
    
    const form = useForm({
        initialValues: {
          username: '',
          password: '',
        },
    
        validate: {
          username: ((value) => value ? null : 'Invalid username'), //validation if needed
          password: ((value)=> value ? null : 'Invalid password') //validation if needed
        },
      });
    
    const [error, setError] = useState('')

    const onSubmit = (values:{username:string,password:string}) => {
        if (signIn(values.username,values.password)) {
            router.push('/dashboard')
        } else {
            setError('Username or password is invalid.')
        }
    }  

    useEffect(() => {
        // redirect to dashboard if already logged in
        if (isLoggedIn()) {
            router.push('/dashboard');
        }
    }, []);
      return (
        <div style={{display:'flex', height:'100vh',flexDirection:'column',alignItems:'center',justifyContent:'center'}}>
            <Text size="xl" weight={500}>iMatrix Dashboard</Text>
            <Card shadow="sm" style={{width:'600px',height:'320px', display:'flex', alignItems:'center',justifyContent:'center'}}>
                    <Box style={{width:'100%',padding:'24px 48px'}}>
                        <form onSubmit={form.onSubmit((values) => onSubmit(values))}>
                            <TextInput
                            required
                            label="Username"
                            style={{marginBottom:'12px'}}
                            {...form.getInputProps('username')}
                            />

                            <TextInput
                            required
                            label="Password"
                            type='password'
                            style={{marginBottom:'24px'}}
                            {...form.getInputProps('password')}
                            error={error}
                            />

                            <Group position="right" mt="md">
                            <Button type="submit">Submit</Button>
                            </Group>
                        </form>
                    </Box>
            </Card>
        </div>
    );
}
 
export default LogIn;