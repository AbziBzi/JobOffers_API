import React, { useEffect, useState } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

const useStyles = makeStyles((theme) => ({
    root: {
        width: '100%',
        maxWidth: 360,
        backgroundColor: theme.palette.background.paper,
    },
}));

function CompanyList() {
    const companies = [
        {
            id: 2,
            name: "Visma",
            size: 5000,
            industry: "Software Engineering",
            headquarters: "Vilnius",
            social_media: [
                "facebook.com/visma",
                "twitter.com/visma"
            ],
            typeID: 3,
            company_type: {
                id: 3,
                name: "Software house"
            },
            user_id: 5,
            company_administrator: {
                id: 5,
                name: "Hank",
                surname: "Hans",
                email: "hhans0@addtoany.com",
                password: "password123",
                role_id: 1,
                role: {
                    id: 0,
                    name: ""
                }
            },
            offices: [
                {
                    id: 1,
                    country: "Lithuania",
                    city: "Kaunas",
                    zip_code: "12345",
                    street: "Mindaugo prospektas",
                    building_nr: "12",
                    company_id: 2
                }
            ]
        },
        {
            id: 2,
            name: "name2"
        }
    ]
    return (
        <List dense className={useStyles.root}>
            {companies.map((value) => {
                return (
                    <ListItem key={value} button>
                        <ListItemText primary={`Line item ${value.id}`} />
                        <ListItemText primary={`Description of item ${value.name}`} />
                    </ListItem>
                );
            })}
        </List>
    );
}

export default CompanyList