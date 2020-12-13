import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import JobOfferList from './components/JobOfferList';
import CompaniesPage from './pages/CompaniesPage';
import CompanyPage from './pages/CompanyPage';
import NavBar from './components/NavBar'
import Footer from './components/Footer'
import './App.css';

function App() {
  return (
    <BrowserRouter>
      <header>
        <NavBar />
      </header>
      <section>
        <Switch>
          <Route path="/companies" exact component={CompaniesPage} />
          <Route path="/companies/:id" exact component={CompanyPage} />
          <Route path="/jobs" exact component={JobOfferList} />
        </Switch>
      </section>
      <footer>
        <Footer />
      </footer>
    </BrowserRouter>
  );
}

export default App;
