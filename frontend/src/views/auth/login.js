import * as React from 'react';
import Button from '@mui/material/Button';
import Container from '@mui/material/Container';
import { Card, FormControl, Input } from '@mui/material';
import GoogleIcon from '@mui/icons-material/Google';
import GitHubIcon from '@mui/icons-material/GitHub';

export default function Login() {
  const handleGoogleLogin = () => {
    fetch('http://localhost:3000/api/auth/google/auth')
        .then(response => response.json())
        .then(data => {
          if (data.code === 200) {
            console.log(data.data)
            window.location.replace(data.data.Url);
          }
        })
  }
  const handleGetGithubAuthURL = () => {
    fetch('http://localhost:3000/api/auth/github/auth')
        .then(response => response.json())
        .then(data => {
          if (data.code === 200) {
            console.log(data.data)
            window.location.replace(data.data.Url);
          }
        })
  }
  return (
    <div>
      <Container fixed>
        <Card style={{textAlign: 'center'}}>
          <div>
            <h1>WellCome to Oauth2 demo web</h1>
          </div>
          <FormControl>
            <h2>Login</h2>
            <div>
              <div>
                <Input id='username' placeholder='*username'></Input>
              </div>
              <div style={{paddingTop: '20px'}}>
                <Input id='password' placeholder='*password' type='password' ></Input>
              </div>
            </div>
            <div style={{marginTop: '20px'}}>             
              <div>
                <Button variant="contained">Login</Button>
              </div> 
              <div style={{paddingTop: '20px'}}>
                <Button onClick={handleGoogleLogin}><GoogleIcon></GoogleIcon></Button>
                <Button onClick={handleGetGithubAuthURL}><GitHubIcon></GitHubIcon></Button>
              </div>
            </div>
          </FormControl>
        </Card>
      </Container>
    </div>
  )
}