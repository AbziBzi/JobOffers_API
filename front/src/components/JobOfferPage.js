import React, { useEffect, useState } from 'react';

function JobOfferPage() {
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
      <ul>
        {jobOffers.map(item => (
          <li key={item.id}>
            {item.name} {item.description}
          </li>
        ))}
      </ul>
    );
  }
}

export default JobOfferPage