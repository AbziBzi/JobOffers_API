import React from 'react';
import Typography from '@material-ui/core/Typography';
import Card from '@material-ui/core/Card';
import { LocationOn, Business } from '@material-ui/icons';
import Chip from '@material-ui/core/Chip';
import Box from '@material-ui/core/Box';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
    offerTitle: {
      display: 'flex',
    },
    secondary: {
      display: 'flex',
      alignItems: 'center'
    },
    newChip: {
      marginLeft: 15,
    },
    experienceChip: {
      maxWidth: 70,
      marginTop: 2,
    }
  }));

function JobOfferCard(props) {
  const classes = useStyles();

  const now = new Date();
  const publicationDate = new Date(props.offer.publication_time);
  const difference = now - publicationDate;
  const publishedDaysAgo = Math.floor((difference / (60*60*24*1000)));

  return (
    <Card>
      <Box display="flex"
           justifyContent="space-between" 
            p={1}>
        <Box display="flex" 
            flexDirection="column">
          <Typography className={classes.offerTitle}
                      component="h5" 
                      variant="h5"
                      color="primary">
            {props.offer.name}
          </Typography>
          <Box className={classes.secondary}>
            <Business fontSize="small"
                      color="disabled"/>
            <Typography variant="subtitle1" color="textSecondary">
              {props.company.name}
            </Typography>
            <LocationOn fontSize="small"
                        color="disabled"/>
            <Typography variant="subtitle1" color="textSecondary">
              {props.company.headquarters}
            </Typography>
          </Box>
        </Box>
        <Box display="flex">
          <Box display="flex"
                flexWrap="wrap-reverse"
                flexDirection="column">
            <Box display="flex"
                  alignItems="center">
              <Typography variant="h5" color="secondary">
                {props.offer.salary_min} - {props.offer.salary_max} €
              </Typography>
              {publishedDaysAgo < 2
                ? <Chip color="primary" 
                        size="small"
                        label="NEW" 
                        className={classes.newChip}/>
                : <Box className={classes.newChip}>{publishedDaysAgo + ' d. ago'}</Box>
              }
            </Box>
            <Chip variant="outlined"
                  size="small"
                  label={props.experience.name}
                  className={classes.experienceChip}/>
          </Box>
        </Box>
      </Box>
    </Card>
  )
}

export default JobOfferCard;