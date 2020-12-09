import React from 'react'
import {Grid, Paper, Container, Card, Typography} from '@material-ui/core'

function JobOfferPage() {
    return (
        <Container maxWidth="md">
            <Grid container 
                    spacing={3}
                    justify="space-around">
                <Grid item xs={12}>
                <Card>
                    <Typography variant="h2" color="primary"> 
                        Junior Java Developer
                    </Typography>
                </Card>
                </Grid>
                <Grid item xs={3}>
                <Paper >xs=6</Paper>
                </Grid>
                <Grid item xs={3}>
                <Paper>xs=6</Paper>
                </Grid>
                <Grid item xs={3}>
                <Paper>xs=3</Paper>
                </Grid>
                <Grid item xs={12}>
                <Paper>xs=3</Paper>
                </Grid>
            </Grid>
        </Container>
    )
}

export default JobOfferPage;