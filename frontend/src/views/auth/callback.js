import * as React from 'react';
import { Navigate, useParams } from 'react-router-dom';

export default function Callback() {
  const { providerID } = useParams()
  const search = window.location.search;
  const query = new URLSearchParams(search);

  switch (providerID) {
    case "google":
      const code = query.get('code')
      fetch('http://localhost:3000/api/auth/google/callback?' + new URLSearchParams({
        code: code,
      }))
      .then(response => response.json())
      .then(data => {
        if (data.code === 200) {
          // TODO: login reply info
          localStorage.setItem('LOGIN_INFO', JSON.stringify(data.data))

          // TODO: redirect
          window.location.replace("/admin/dashboard");
        }else {
          <Navigate to="/user/login" replace></Navigate>
        }
      })
      break;
    default:
      console.log("unspecified")
      break;
  }
  return (
    <></>
  )
}