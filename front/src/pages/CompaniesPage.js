import React, { useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles'
import {
    Grid,
    CircularProgress
} from '@material-ui/core/'
import CompanyCard from '../components/CompanyCard'

const useStyles = makeStyles(theme => ({
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

function CompaniesPage() {
    const classes = useStyles()
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [companies, setCompanies] = useState([]);

    useEffect(() => {
        fetch("http://localhost:3033/api/companies", {
            method: 'GET',
            mode: 'cors',
            headers: {
                'Content-Type': 'application/json'
            },
        })
            .then(res => res.json())
            .then(
                (result) => {
                    setIsLoaded(true);
                    setCompanies(result);
                },
                (error) => {
                    setIsLoaded(true);
                    setError(error);
                }
            )
    }, [])

    if (error) {
        return <div>Error: {error.message}</div>;
    } else if (!isLoaded) {
        return (
            <div className={classes.spinner}>
                <CircularProgress size={100} />
            </div>
        )
    } else {
        return (
            <div className={classes.root}>
                <Grid
                    container
                    spacing={2}
                    direction="row"
                    justify="flex-start"
                    alignItems="flex-start">
                    {companies.map(company => (
                        <Grid item xs={6} key={companies.indexOf(company)}>
                            <CompanyCard company={company} />
                        </Grid>
                    ))}
                </Grid>
            </div>
        )
    }
}

export default CompaniesPage