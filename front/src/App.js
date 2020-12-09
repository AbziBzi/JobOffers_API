import React from 'react';
import './App.css';
import Container from '@material-ui/core/Container';
import JobOfferPage from './components/JobOfferPage';
import CompaniesPage from './pages/CompaniesPage';

function App() {
  return (
    <div className="App">
      <Container maxWidth="md" >
        <CompaniesPage />
      </Container>
    </div>
  );
}

export default App;
