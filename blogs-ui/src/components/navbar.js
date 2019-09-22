import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Button from '@material-ui/core/Button';
import DeveloperModeIcon from '@material-ui/icons/DeveloperMode';

const useStyles = makeStyles(theme => ({
    bar: {
        justifyContent: "space-between",
    },
    icon: {
        fontSize: 50,
        marginLeft: 50,
    },
    button: {
        marginRight: 50,
    },
}));

export default function ButtonAppBar() {
    const classes = useStyles();

    return (
        <div>
            <AppBar position="static" color="grey.300">
                <Toolbar className={classes.bar}>
                    <DeveloperModeIcon className={classes.icon}/>
                    <Button color="inherit" className={classes.button}>Login</Button>
                </Toolbar>
            </AppBar>
        </div>
    );
}