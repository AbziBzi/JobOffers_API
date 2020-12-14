import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import JobOfferList from './components/JobOfferList';
import CompaniesPage from './pages/CompaniesPage';
import CompanyPage from './pages/CompanyPage';
import NavBar from './components/NavBar'
import SignInPage from './pages/SignInPage'
import SignUpPage from './pages/SignUpPage'
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
          <Route path="/login" exact component={SignInPage} />
          <Route path="/register" exact component={SignUpPage} />
        </Switch>
      </section>
      <footer>
        <Footer />
      </footer>
    </BrowserRouter>
  );
}

export default App;
