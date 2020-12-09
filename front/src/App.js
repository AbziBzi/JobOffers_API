import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import JobOfferPage from './components/JobOfferPage';
import CompaniesPage from './pages/CompaniesPage';
import CompanyPage from './pages/CompanyPage';
import NavBar from './components/NavBar'

function App() {
  return (
    <div>
      <BrowserRouter>
        <NavBar />
        <Switch>
          <Route path="/companies" exact component={CompaniesPage} />
          <Route path="/companies/:id" exact component={CompanyPage} />
          <Route path="/jobs" exact component={JobOfferPage} />
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
