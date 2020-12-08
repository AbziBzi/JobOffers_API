import React from 'react';
import './App.css';
import Container from '@material-ui/core/Container';
import SignUpPage from './components/SignUpPage'

function App() {
  return (
    <div className="App">
      <Container maxWidth="md" >
        <SignUpPage />
      </Container>
    </div>
  );
}

export default App;
