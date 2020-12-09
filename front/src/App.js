import React from 'react';
import './App.css';
import Container from '@material-ui/core/Container';
import SignUpPage from './components/SignUpPage'
import JobOfferPage from './components/JobOfferPage';

function App() {
  return (
    <div className="App">
      <Container maxWidth="md" >
        <JobOfferPage />
      </Container>
    </div>
  );
}

export default App;
