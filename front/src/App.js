import React from 'react';
import './App.css';
import Container from '@material-ui/core/Container';
import SignUpPage from './pages/SignUpPage'
import JobOfferCard from './components/JobOfferCard'

function App() {
  return (
    <div className="App">
      <Container maxWidth="md" >
        <CompanyCard />
      </Container>
    </div>
  );
}

export default App;
