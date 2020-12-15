import React, { useEffect, useState } from 'react';
import { Grid, Paper, Container, Card, Typography, Box, Button, CircularProgress } from '@material-ui/core'
import Business from '@material-ui/icons/Business';
import TrendingUpIcon from '@material-ui/icons/TrendingUp';
import EuroIcon from '@material-ui/icons/Euro';
import Send from '@material-ui/icons/Send';
import { makeStyles } from '@material-ui/core/styles';
import { Dialog, DialogActions, DialogContent, DialogTitle, TextField } from '@material-ui/core';

const useStyles = makeStyles(() => ({
    paperPadding: {
        paddingTop: 20,
        paddingBottom: 20
    }
}));

function JobOfferPage(props) {
    const classes = useStyles();
    const jobID = props.match.params.id
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [job, setJob] = useState([]);

    useEffect(() => {
        fetch(`http://localhost:3033/api/jobOffers/${jobID}`, {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'
            },
        })
            .then(res => res.json())
            .then(
                (result) => {
                    setJob(result);
                    if (result.error != null) {
                        setError(result.error);
                        console.log(result)
                        setIsLoaded(true);
                    }
                    setIsLoaded(true);
                }
            )
    }, [])

    const [open, setOpen] = React.useState(false);

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    if (error != null) {
        return <div>Error: {error}</div>;
    } else if (!isLoaded) {
        return (
            <div className={classes.spinner}>
                <CircularProgress size={100} />
            </div>
        )
    } else {
        const offer = job;
        const company = job.company;
        const experience = job.experience;
        return (
            <Container maxWidth="md">
                <Box className={classes.paperPadding}>
                    <Paper>
                        <Box p={5}>
                            <Grid container
                                spacing={3}
                                justify="space-between">
                                <Grid item xs={12}>
                                    <Card>
                                        <Box p={1}>
                                            <Typography variant="h2" color="primary">
                                                {job.name}
                                            </Typography>
                                        </Box>
                                    </Card>
                                </Grid>
                                <Grid item xs={4}>
                                    <Card>
                                        <Box p={1}
                                            display="flex"
                                            flexDirection="column"
                                            justifyContent="space-evenly"
                                            alignItems="center">
                                            <Business fontSize="small"
                                                color="disabled" />
                                            <Typography variant="h6">
                                                {job.company.name} | {job.company.headquarters}
                                            </Typography>
                                            <Typography variant="subtitle2">
                                                Company | Location
                                        </Typography>
                                        </Box>
                                    </Card>
                                </Grid>
                                <Grid item xs={4}>
                                    <Card>
                                        <Box p={1}
                                            display="flex"
                                            flexDirection="column"
                                            justifyContent="space-evenly"
                                            alignItems="center">
                                            <TrendingUpIcon fontSize="small"
                                                color="disabled" />
                                            <Typography variant="h6"
                                                align="center">
                                                {job.experience.name}
                                            </Typography>
                                            <Typography variant="subtitle2">
                                                Experience
                                        </Typography>
                                        </Box>
                                    </Card>
                                </Grid>
                                <Grid item xs={4}>
                                    <Card>
                                        <Box p={1}
                                            display="flex"
                                            flexDirection="column"
                                            justifyContent="space-evenly"
                                            alignItems="center">
                                            <EuroIcon fontSize="small"
                                                color="disabled" />
                                            <Typography variant="h6"
                                                align="center"
                                                color="secondary">
                                                {job.salary_min} - {job.salary_max} €
                                        </Typography>
                                            <Typography variant="subtitle2">
                                                Salary
                                        </Typography>
                                        </Box>
                                    </Card>
                                </Grid>
                                <Grid item xs={12}>
                                    <Card>
                                        <Box p={1}>
                                            <Typography variant="h4">
                                                Description
                                        </Typography>
                                            <hr></hr>
                                            <Typography variant="h6">
                                                {job.description}
                                            </Typography>
                                        </Box>
                                    </Card>
                                </Grid>
                                <Grid item xs={10}>
                                </Grid>
                                <Grid item xs={2}>
                                    <Button variant="contained"
                                        color="secondary"
                                        size="large"
                                        endIcon={<Send />}
                                        onClick={handleClickOpen}>
                                        Apply
                                </Button>
                                    <Dialog open={open} onClose={handleClose} aria-labelledby="form-dialog-title">
                                        <DialogTitle id="form-dialog-title">Apply</DialogTitle>
                                        <DialogContent>
                                            <Grid container spacing={2}>
                                                <Grid item xs={6}>
                                                    <TextField
                                                        margin="dense"
                                                        id="first-name"
                                                        label="First Name"
                                                        type="text"
                                                        fullWidth
                                                    />
                                                </Grid>
                                                <Grid item xs={6}>
                                                    <TextField
                                                        margin="dense"
                                                        id="last-name"
                                                        label="Last Name"
                                                        type="text"
                                                        fullWidth
                                                    />
                                                </Grid>
                                                <Grid item xs={12}>
                                                    <TextField
                                                        margin="dense"
                                                        id="email"
                                                        label="Email Address"
                                                        type="email"
                                                        fullWidth
                                                    />
                                                </Grid>
                                                <Grid item xs={12}>
                                                    <TextField
                                                        multiline
                                                        rows={4}
                                                        margin="dense"
                                                        id="cover"
                                                        label="Cover Letter"
                                                        type="text"
                                                        fullWidth
                                                    />
                                                </Grid>
                                                <Grid item xs={12}>
                                                    <Button variant="contained"
                                                        component="label">
                                                        Upload CV
                                                        <input type="file"
                                                            hidden />
                                                    </Button>
                                                </Grid>
                                            </Grid>
                                        </DialogContent>
                                        <DialogActions>
                                            <Button onClick={handleClose} color="primary">
                                                Cancel
                                        </Button>
                                            <Button onClick={handleClose} color="primary">
                                                Send
                                        </Button>
                                        </DialogActions>
                                    </Dialog>
                                </Grid>
                            </Grid>
                        </Box>
                    </Paper>
                </Box>
            </Container>
        )
    }
}

export default JobOfferPage;