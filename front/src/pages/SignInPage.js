import React, { useState, useContext } from 'react';
import { useHistory } from "react-router-dom";
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Link from '@material-ui/core/Link';
import Grid from '@material-ui/core/Grid';
import Typography from '@material-ui/core/Typography';
import { makeStyles } from '@material-ui/core/styles';
import Container from '@material-ui/core/Container';
import { UserContext } from '../UserContext';

const useStyles = makeStyles((theme) => ({
    formControl: {
        marginTop: theme.spacing(2),
        minWidth: 150
    },
    paper: {
        marginTop: theme.spacing(8),
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    avatar: {
        margin: theme.spacing(1),
        backgroundColor: theme.palette.secondary.main,
    },
    form: {
        width: '100%', // Fix IE 11 issue.
        marginTop: theme.spacing(1),
    },
    submit: {
        margin: theme.spacing(3, 0, 2),
    },
}));


function SignInPage() {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [token, setToken] = useState("")
    const [isLogged, setLogged] = useState(false)
    const [error, setError] = useState("")
    const classes = useStyles();
    const user = useContext(UserContext)
    const history = useHistory();

    const onLogin = () => {
        const bodyJSON = JSON.stringify({ email: email, password: password })
        setLogged(false)
        fetch("http://localhost:3033/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Accept: "application/json"
            },
            body: bodyJSON
        }).then(response => response.json())
            .then(jsonResponse => {
                setToken(jsonResponse)
                if (jsonResponse.error != null) {
                    setError(jsonResponse.error);
                    console.log(jsonResponse)
                }
                setLogged(true)
            })
            .then(onGetUserData())
            .catch(error => console.log(error))
    }

    const z = () => {
        if (isLogged == true) {
            fetch("http://localhost:3033/api/users/token", {
                method: "GET",
                mode: 'cors',
                headers: {
                    'Authorization': 'Bearer ' + token,
                    'Content-Type': 'application/json'
                }
            }).then(response => response.json())
                .then(jsonResponse => {
                    user.setToken(token)
                    user.setId(jsonResponse.id)
                    user.setRoleId(jsonResponse.role_id)
                    console.log(user)
                    history.push('/jobs')
                })
                .catch(error => console.log(error))
        } else {
            console.log("Wrong email or password")
        }
    }

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <div className={classes.paper}>
                <Link href="https://www.google.com/" variant="h4">
                    {'SiteName'}
                </Link>
                <Typography component="h1" variant="h5">
                    Sign in
                </Typography>
                {(error == null) && <Typography component="h1" variant="h6" color="error">
                    Wrong email or password.
                </Typography>}
                <TextField
                    variant="outlined"
                    margin="normal"
                    required
                    fullWidth
                    id="email"
                    label="Email Address"
                    name="email"
                    autoComplete="email"
                    autoFocus
                    onChange={e => setEmail(e.target.value)}
                />
                <TextField
                    variant="outlined"
                    margin="normal"
                    required
                    fullWidth
                    name="password"
                    label="Password"
                    type="password"
                    id="password"
                    autoComplete="current-password"
                    onChange={e => setPassword(e.target.value)}
                />
                <Button
                    type="submit"
                    fullWidth
                    variant="contained"
                    color="primary"
                    className={classes.submit}
                    onClick={onLogin}
                >
                    Sign In
                    </Button>
                <Grid container>
                    <Grid item>
                        <Link href="/register" variant="body2">
                            {"Don't have an account? Sign Up"}
                        </Link>
                    </Grid>
                </Grid>
            </div>
        </Container>
    );
}

export default SignInPage