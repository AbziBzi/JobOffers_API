import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import JobOfferList from './components/JobOfferList';
import CompaniesPage from './pages/CompaniesPage';
import CompanyPage from './pages/CompanyPage';
import NavBar from './components/NavBar'
import SignInPage from './pages/SignInPage'
import Footer from './components/Footer'
import './App.css';
import {createMuiTheme, ThemeProvider} from '@material-ui/core'
import JobOfferPage from './pages/JobOfferPage'

const theme = createMuiTheme({
  typography: {
    fontFamily: [
      'Heebo',
      'serif',
    ].join(','),
},});

function App() {
  return (
    <BrowserRouter>
      <ThemeProvider theme={theme}>
        <header>
          <NavBar />
        </header>
        <section>
          <Switch>
            <Route path="/companies" exact component={CompaniesPage} />
            <Route path="/companies/:id" exact component={CompanyPage} />
            <Route path="/jobs" exact component={JobOfferList} />
            <Route path="/jobs/:id" exact component={JobOfferPage} />
            <Route path="/login" exact component={SignInPage} />
          </Switch>
          {/* <JobOfferPage /> */}
        </section>
        <footer>
          <Footer />
        </footer>
      </ThemeProvider>
    </BrowserRouter>
  );
}

export default App;
