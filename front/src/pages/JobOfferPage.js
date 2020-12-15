import React from 'react'
import {Grid, Paper, Container, Card, Typography, Box, Button} from '@material-ui/core'
import Business from '@material-ui/icons/Business';
import TrendingUpIcon from '@material-ui/icons/TrendingUp';
import EuroIcon from '@material-ui/icons/Euro';
import Send from '@material-ui/icons/Send';
import { makeStyles } from '@material-ui/core/styles';
import {Dialog, DialogActions, DialogContent, DialogTitle, TextField} from '@material-ui/core';

const useStyles = makeStyles(() => ({
    paperPadding: {
      paddingTop: 20,
      paddingBottom: 20
    }
  }));

function JobOfferPage(props) {
    const classes = useStyles();

    const [open, setOpen] = React.useState(false);

    const handleClickOpen = () => {
        setOpen(true);
    };

    const handleClose = () => {
        setOpen(false);
    };

    const offer = props.location.state.offer;
    const company = props.location.state.company;
    const experience = props.location.state.experience;
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
                                        {offer.name}
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
                                                color="disabled"/>
                                        <Typography variant="h6"> 
                                            {company.name} | {company.headquarters}
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
                                                        color="disabled"/>
                                        <Typography variant="h6"
                                                    align="center"> 
                                            {experience.name}
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
                                                  color="disabled"/>
                                        <Typography variant="h6"
                                                    align="center"
                                                    color="secondary"> 
                                            {offer.salary_min} - {offer.salary_max} €
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
                                            {offer.description}
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
                                        endIcon={<Send/>}
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
                                                               hidden/>
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

export default JobOfferPage;