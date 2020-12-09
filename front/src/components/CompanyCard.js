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

function CompanyCard(props) {
    const company = props.company

    const handleClick = () => {
        console.log("one day it will work and in this day you will know that every click you do...")
    }
    return (
        <Card>
            <CardActionArea onClick={handleClick}>
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
    );
}

export default CompanyCard