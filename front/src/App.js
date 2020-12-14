import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import JobOfferList from './components/JobOfferList';
import CompaniesPage from './pages/CompaniesPage';
import CompanyPage from './pages/CompanyPage';
import NavBar from './components/NavBar'
<<<<<<< HEAD
import SignInPage from './pages/SignInPage'
=======
import Footer from './components/Footer'
import './App.css';
>>>>>>> 50c9f18760206096647de3977593ed3aae37e713

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
        </Switch>
      </section>
      <footer>
        <Footer />
      </footer>
    </BrowserRouter>
  );
}

export default App;
