import React, { useState, useContext } from 'react';
import { useHistory } from "react-router-dom";
import { makeStyles, useTheme } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import { UserContext } from '../UserContext';
import IconButton from '@material-ui/core/IconButton';
import Drawer from '@material-ui/core/Drawer';
import Toolbar from '@material-ui/core/Toolbar';
import List from '@material-ui/core/List';
import Divider from '@material-ui/core/Divider';
import MenuIcon from '@material-ui/icons/Menu';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
    root: {
        display: 'flex',
    },
    appBar: {
        transition: theme.transitions.create(['margin', 'width'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
        }),
    },
    appBarShift: {
        width: `calc(100% - ${drawerWidth}px)`,
        marginLeft: drawerWidth,
        transition: theme.transitions.create(['margin', 'width'], {
            easing: theme.transitions.easing.easeOut,
            duration: theme.transitions.duration.enteringScreen,
        }),
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    hide: {
        display: 'none',
    },
    drawer: {
        width: drawerWidth,
        flexShrink: 0,
    },
    drawerPaper: {
        width: drawerWidth,
    },
    drawerHeader: {
        display: 'flex',
        alignItems: 'center',
        padding: theme.spacing(0, 1),
        // necessary for content to be below app bar
        ...theme.mixins.toolbar,
        justifyContent: 'flex-end',
    },
    content: {
        flexGrow: 1,
        padding: theme.spacing(3),
        transition: theme.transitions.create('margin', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
        }),
        marginLeft: -drawerWidth,
    },
    contentShift: {
        transition: theme.transitions.create('margin', {
            easing: theme.transitions.easing.easeOut,
            duration: theme.transitions.duration.enteringScreen,
        }),
        marginLeft: 0,
    },
}));

export default function SimpleTabs() {
    const user = useContext(UserContext)
    const classes = useStyles();
    const [value, setValue] = React.useState(0);
    const history = useHistory();
    const [width, setWidth] = React.useState(window.innerWidth);
    const breakpoint = 900;
    const [anchorEl, setAnchorEl] = React.useState(null);
    const theme = useTheme();
    const [open, setOpen] = React.useState(false);

    const handleDrawerOpen = () => {
        setOpen(true);
    };

    const handleDrawerClose = () => {
        setOpen(false);
    };

    React.useEffect(() => {
        const handleWindowResize = () => setWidth(window.innerWidth)
        window.addEventListener("resize", handleWindowResize);

        // Return a function from the effect that removes the event listener
        return () => window.removeEventListener("resize", handleWindowResize);
    }, []);
    const handleChange = (event, newValue) => {
        setValue(newValue);
    };

    function onLogOut() {
        user.setToken("")
        user.setId(0)
        user.setRoleId(0)
        history.push("/login")
    }
    if (width > breakpoint) {
        if (user.roleId == 1) {
            console.log(user.roleId)
            return (
                <div className={classes.root}>
                    <AppBar position="static">
                        <Tabs value={value} onChange={handleChange} aria-label="simple tabs example">
                            <Tab label="Open Jobs" href="/jobs" />
                            <Tab label="Companies" href="/companies" />
                            <Tab label="Log Out" onClick={onLogOut} />
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
                            <Tab label="Log Out" onClick={onLogOut} />
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
    } else {
        if (user.roleId == 1) {
            console.log(user.roleId)
            return (
                <div className={classes.root}>
                    <AppBar position="static">
                        <Toolbar>
                            <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu" onClick={handleDrawerOpen}>
                                <MenuIcon />
                            </IconButton>
                            <Drawer
                                className={classes.drawer}
                                variant="persistent"
                                anchor="left"
                                open={open}
                                classes={{
                                    paper: classes.drawerPaper,
                                }}
                            >
                                <div className={classes.drawerHeader}>
                                    <IconButton onClick={handleDrawerClose}>
                                        {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
                                    </IconButton>
                                </div>
                                <Divider />
                                <List>
                                    <ListItem button onClick={() => { history.push("/jobs") }}>
                                        <ListItemText primary="Open Jobs" />
                                    </ListItem>
                                    <ListItem button onClick={() => { history.push("/companies") }}>
                                        <ListItemText primary="Companies" />
                                    </ListItem>
                                    <ListItem button onClick={onLogOut}>
                                        <ListItemText primary="Log Out" />
                                    </ListItem>
                                </List>
                            </Drawer>
                        </Toolbar>
                    </AppBar>
                </div>
            );
        }
        else if (user.roleId == 2) {
            console.log(user.roleId)
            return (
                <div className={classes.root}>
                    <AppBar position="static">
                        <Toolbar>
                            <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu" onClick={handleDrawerOpen}>
                                <MenuIcon />
                            </IconButton>
                            <Drawer
                                className={classes.drawer}
                                variant="persistent"
                                anchor="left"
                                open={open}
                                classes={{
                                    paper: classes.drawerPaper,
                                }}
                            >
                                <div className={classes.drawerHeader}>
                                    <IconButton onClick={handleDrawerClose}>
                                        {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
                                    </IconButton>
                                </div>
                                <Divider />
                                <List>
                                    <ListItem button onClick={() => { history.push("/jobs") }}>
                                        <ListItemText primary="Open Jobs" />
                                    </ListItem>
                                    <ListItem button onClick={() => { history.push("/companies") }}>
                                        <ListItemText primary="Companies" />
                                    </ListItem>
                                    <ListItem button >
                                        <ListItemText primary="Add Job Offer" />
                                    </ListItem>
                                    <ListItem button onClick={onLogOut}>
                                        <ListItemText primary="Log Out" />
                                    </ListItem>
                                </List>
                            </Drawer>
                        </Toolbar>
                    </AppBar>
                </div>
            );
        }
        else {
            console.log(user.roleId)
            return (
                <div className={classes.root}>
                    <AppBar position="static">
                        <Toolbar>
                            <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu" onClick={handleDrawerOpen}>
                                <MenuIcon />
                            </IconButton>
                            <Drawer
                                className={classes.drawer}
                                variant="persistent"
                                anchor="left"
                                open={open}
                                classes={{
                                    paper: classes.drawerPaper,
                                }}
                            >
                                <div className={classes.drawerHeader}>
                                    <IconButton onClick={handleDrawerClose}>
                                        {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
                                    </IconButton>
                                </div>
                                <Divider />
                                <List>
                                    <ListItem button onClick={() => { history.push("/jobs") }}>
                                        <ListItemText primary="Open Jobs" />
                                    </ListItem>
                                    <ListItem button onClick={() => { history.push("/companies") }}>
                                        <ListItemText primary="Companies" />
                                    </ListItem>
                                    <ListItem button onClick={() => { history.push("/login") }}>
                                        <ListItemText primary="Sign In" />
                                    </ListItem>
                                    <ListItem button onClick={() => { history.push("/signup") }}>
                                        <ListItemText primary="Sign Up" />
                                    </ListItem>
                                </List>
                            </Drawer>
                        </Toolbar>
                    </AppBar>
                </div>
            );
        }
    }
}