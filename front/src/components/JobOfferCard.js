import React from 'react';
import Typography from '@material-ui/core/Typography';
import Card from '@material-ui/core/Card';
import { LocationOn, Business } from '@material-ui/icons';
import Chip from '@material-ui/core/Chip';
import Box from '@material-ui/core/Box';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
    cardContent: {
      display: 'flex',
      justifyContent: 'space-between',
    },
    secondary: {
      display: 'flex',
      alignItems: 'center'
    },
    newChip: {
      marginLeft: 15,
    },
    compatenceChip: {
      maxWidth: 70,
      marginTop: 2,
    }
  }));

function JobOfferCard() {
  const classes = useStyles();

  return (
    <Card>
      <Box className={classes.cardContent} 
            p={1}>
        <Box display="flex" 
            flexDirection="column">
          <Typography component="h5" variant="h5">
            Junior Java Developer
          </Typography>
          <Box className={classes.secondary}>
            <Business fontSize="small"
                      color="disabled"/>
            <Typography variant="subtitle1" color="textSecondary">
              Visma
            </Typography>
            <LocationOn fontSize="small"
                        color="disabled"/>
            <Typography variant="subtitle1" color="textSecondary">
              Vilnius
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
                6000 - 12000 PLN
              </Typography>
              <Chip color="primary" 
                    size="small"
                    label="NEW" 
                    className={classes.newChip}/>
            </Box>
            <Chip variant="outlined"
                  size="small"
                  label="Junior"
                  className={classes.compatenceChip}/>
          </Box>
        </Box>
      </Box>
    </Card>
  )
}

export default JobOfferCard;