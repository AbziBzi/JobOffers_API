import React from 'react'
import { Box, Typography, Link } from '@material-ui/core'

function Footer() {
    return (
        <Box bgcolor="lightgray"
            display="flex"
            justifyContent="center"
            alignItems="center"
            padding="20px">
            <Typography>
                <Link href="https://github.com/AbziBzi"
                    rel="noopener">
                    Made by Abzibzi © 2020
                </Link>
            </Typography>
        </Box>
    );
}

export default Footer;