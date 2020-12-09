import React, { useEffect, useState } from 'react';
import JobOfferCard from './JobOfferCard';
import Box from '@material-ui/core/Box';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container'

const useStyles = makeStyles(() => ({
  jobOfferCard: {
    marginTop: 10
  }
}));

function JobOfferList() {
    const classes = useStyles();

    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [jobOffers, setJobOffers] = useState([]);

  useEffect(() => {
    fetch("http://3.124.191.230:3030/api/jobOffers", {
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
          setJobOffers(result);
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
    return <div>Loading...</div>;
  } else {
    return (
      <Container maxWidth="md">
        <Box display="flex"
            flexDirection="column">
          {jobOffers.map(item => (
            <div key={item.id}
            className={classes.jobOfferCard}>
              <JobOfferCard offer={item}
                            company={item.company}
                            experience={item.experience}/>
            </div>
          ))}
        </Box>
      </Container>
    );
  }
}

export default JobOfferList