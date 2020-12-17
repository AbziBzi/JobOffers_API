import React, { useState, useContext } from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import Typography from '@material-ui/core/Typography';
import Box from '@material-ui/core/Box';
import { UserContext } from '../UserContext';


const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
        backgroundColor: theme.palette.background.paper,
    },
}));

export default function SimpleTabs() {
    const user = useContext(UserContext)
    const classes = useStyles();
    const [value, setValue] = React.useState(0);

    const handleChange = (event, newValue) => {
        setValue(newValue);
    };

    if (user.roleId == 1) {
        console.log(user.roleId)
        return (
            <div className={classes.root}>
                <AppBar position="static">
                    <Tabs value={value} onChange={handleChange} aria-label="simple tabs example">
                        <Tab label="Open Jobs" href="/jobs" />
                        <Tab label="Companies" href="/companies" />
                        <Tab label="Log Out" />
                    </Tabs>
                </AppBar>
            </div>
        );
    }
    else if (user.role == 2) {
        console.log(user.roleId)
        return (
            <div className={classes.root}>
                <AppBar position="static">
                    <Tabs value={value} onChange={handleChange} aria-label="simple tabs example">
                        <Tab label="Open Jobs" href="/jobs" />
                        <Tab label="Companies" href="/companies" />
                        <Tab label="Add Job Offer" />
                        <Tab label="Log Out" />
                    </Tabs>
                </AppBar>
            </div>
        );
    }
    else {
        console.log(user.roleId)
        return (
            <div className={classes.root}>
                <AppBar position="static">
                    <Tabs value={value} onChange={handleChange} aria-label="simple tabs example">
                        <Tab label="Open Jobs" href="/jobs" />
                        <Tab label="Companies" href="/companies" />
                        <Tab label="Sign In" href="/login" />
                        <Tab label="Sign Up" href="/signup" />
                    </Tabs>
                </AppBar>
            </div>
        );
    }
}