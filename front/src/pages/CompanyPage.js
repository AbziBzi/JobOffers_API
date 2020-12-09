import React from 'react';
import { BrowserRouter, Switch, Route, Link } from 'react-router-dom';
import { makeStyles } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import AppBar from '@material-ui/core/AppBar';
import CssBaseline from '@material-ui/core/CssBaseline';
import Toolbar from '@material-ui/core/Toolbar';
import List from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import InboxIcon from '@material-ui/icons/MoveToInbox';
import MailIcon from '@material-ui/icons/Mail';
import DetailsIcon from '@material-ui/icons/Details'
import HomeWorkIcon from '@material-ui/icons/HomeWork';
import WorkOutlineIcon from '@material-ui/icons/WorkOutline';
import SignInPage from './SignInPage'
import CompaniesPage from './CompaniesPage';
import CompanyDetails from '../components/CompanyDetails'
import CompanyOffices from '../components/CompanyOffices'
import CompanyJobs from '../components/CompanyJobs'

const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
    },
    appBar: {
        zIndex: theme.zIndex.drawer + 1,
    },
    drawer: {
        width: 220,
    },
    drawerPaper: {
        width: 220,
        top: 65
    },
    drawerContainer: {
        overflow: 'auto',
    },
    content: {
        flexGrow: 1,
        padding: theme.spacing(3),
    },
}));

function CompanyPage(props) {
    const company = props.location.state.company
    const classes = useStyles();
    return (
        <BrowserRouter>
            <div className={classes.root}>
                <Drawer
                    className={classes.drawer}
                    variant="permanent"
                    classes={{
                        paper: classes.drawerPaper,
                    }}>
                    <div className={classes.drawerContainer}>
                        <List>
                            <ListItem component={Link} to={`/companies/${company.id}`}>
                                <ListItemIcon>
                                    <DetailsIcon />
                                </ListItemIcon>
                                <ListItemText primary="Details" />
                            </ListItem>
                            <ListItem component={Link} to={`/companies/${company.id}/offices`}>
                                <ListItemIcon>
                                    <HomeWorkIcon />
                                </ListItemIcon>
                                <ListItemText primary="Offices" />
                            </ListItem>
                            <ListItem component={Link} to={`/companies/${company.id}/jobs`}>
                                <ListItemIcon>
                                    <WorkOutlineIcon />
                                </ListItemIcon>
                                <ListItemText primary="Open Jobs" />
                            </ListItem>
                        </List>
                    </div>
                </Drawer>
                <main className={classes.content}>
                    <Switch>
                        <Route exact path="/companies/:id" render={() => <CompanyDetails company={company} />} />
                        <Route path="/companies/:id/offices" render={() => <CompanyOffices offices={company.offices} />} />
                        <Route path="/companies/:id/jobs" render={() => <CompanyJobs company={company} />} />
                    </Switch>
                </main>
            </div>
        </BrowserRouter>
    );
}

export default CompanyPage