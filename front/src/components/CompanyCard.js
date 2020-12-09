import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { red } from '@material-ui/core/colors';
import {
    Grid,
    Card,
    CardContent,
    Typography,
    CardHeader,
    CardActionArea
} from '@material-ui/core/'
import { Link } from 'react-router-dom';

const useStyles = makeStyles(theme => ({
    link: {
        textDecoration: 'none'
    },
    root: {
        flexGrow: 1,
        padding: theme.spacing(2)
    },
    spinner: {
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh"
    }
}))

function CompanyCard(props) {
    const classes = useStyles()
    const company = props.company

    const handleClick = () => {
        console.log("one day it will work and in this day you will know that every click you do...")
    }
    return (
        <Link className={classes.link} to={{ pathname: `/companies/${company.id}`, state: { company: company } }} key={company.id}>
            <Card>
                <CardActionArea>
                    <CardHeader
                        title={company.name}
                        subheader={company.industry} />
                    <CardContent>
                        <Grid container
                            spacing={2}
                            direction="row"
                            justify="space-between"
                            alignItems="flex-start"
                        >
                            <Typography variant="body2" color="textSecondary">
                                <b>Headquarter:</b> {company.headquarters}
                            </Typography>
                            <Typography variant="body2" color="textSecondary">
                                <b>Company Type:</b> {company.company_type.name}
                            </Typography>
                        </Grid>
                    </CardContent>
                </CardActionArea>
            </Card>
        </Link>
    );
}

export default CompanyCard