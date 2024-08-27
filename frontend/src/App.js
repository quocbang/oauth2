import './App.css';
import Login from './views/auth/login';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import Callback from './views/auth/callback';
import TopAppBar from './views/components/app_bar/app_bar';
import Dashboard from './views/dashboard/dashboard';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
});

function App() {
  const ProtectedRoute = ({
    redirectPath = '/user/login',
    children,
  }) => {
    const loginInfo = localStorage.getItem('LOGIN_INFO')
    console.log(loginInfo)
    if (!loginInfo) {
      return <Navigate to={redirectPath} replace />;
    }
  
    return children;
  };

  return (
    <ThemeProvider theme={darkTheme}>
      <CssBaseline />
      <BrowserRouter>
        <Routes>
          <Route path="/user/login" element={<Login></Login>}></Route>
          <Route path="/auth/:providerID/callback" element={<Callback />} />
          <Route path="/" element={<ProtectedRoute> <TopAppBar></TopAppBar> </ProtectedRoute>}>
            <Route path="/" element={<Dashboard></Dashboard>} />
          </Route>
          <Route path="*" element={<p>There's nothing here: 404!</p>} />
        </Routes>
      </BrowserRouter>
    </ThemeProvider>
  );
}

export default App;
