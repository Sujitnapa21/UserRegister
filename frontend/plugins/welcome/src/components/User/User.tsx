import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import MenuIcon from '@material-ui/icons/Menu';
import AccountCircle from '@material-ui/icons/AccountCircle';
import Swal from 'sweetalert2'; // alert
import {
  AppBar,
  Toolbar,
  Typography,
  IconButton,
  Grid,
  TextField,
  Container,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  Button
} from '@material-ui/core';

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
    width: '25ch',
    color: "blue"
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 150,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  buttonSty: {
    margin: 'auto',
    display: 'block',
    maxWidth: '100%',
    maxHeight: '100%',
    marginBottom: 50
  }

}));

import { DefaultApi } from '../../api/'; // Api Gennerate From Command

import { EntTitle } from '../../api/models/EntTitle'; //import interface Title
import { EntSex } from '../../api/models/EntSex'; //import interface Sex  
import { EntRole } from '../../api/models/EntRole'; //import interface Role

interface User {
  email: string;
  password: string;
  title: number;
  name: string;
  sex: number;
  age: string;
  birthday: string;
  tel: string;
  role: number;
}


const User1: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  const [user, setUser] = React.useState<Partial<User>>({});

  const [titles, setTitles] = React.useState<EntTitle[]>([]);
  const [sexs, setSexs] = React.useState<EntSex[]>([]);
  const [roles, setRoles] = React.useState<EntRole[]>([]);

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getTitle = async () => {
    const res = await api.listTitle({ limit: 3, offset: 0 });
    setTitles(res);
  };

  const getSex = async () => {
    const res = await api.listSex({ limit: 2, offset: 0 });
    setSexs(res);
  };

  const getRole = async () => {
    const res = await api.listRole({ limit: 2, offset: 0 });
    setRoles(res);
  };


  // Lifecycle Hooks
  useEffect(() => {
    getTitle();
    getSex();
    getRole();
  }, []);

  // set data to object user
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof user;
    const { value } = event.target;
    setUser({ ...user, [name]: value });
    console.log(user);
  };

  // clear input form
  function clear() {
    setUser({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/users';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(user),
    };

    console.log(user); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => {
        console.log(response)
        response.json()
        if (response.ok === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      })
  }

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            ระบบสถานกีฬา
          </Typography>
          <div>
            <IconButton
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              color="inherit"
            >
              <AccountCircle />
            </IconButton>
          </div>
        </Toolbar>
      </AppBar>
      <Container maxWidth="sm">

        <Grid item xs={10}>
          <h2 > ระบบสมัครสมาชิก </h2>
        </Grid>

        <Grid item xs={10}>

        </Grid>

        <Grid container spacing={3}>

          <Grid item xs={6}>
            <TextField
              id="email"
              name="email"
              type="string"
              label="Email"
              variant="outlined"
              fullWidth
              multiline
              value={user.email || ""}
              onChange={handleChange}
            />
          </Grid>

          <Grid item xs={6}>
            <TextField
              id="password"
              name="password"
              type="string"
              autoComplete="current-password"
              label="Password"
              variant="outlined"
              fullWidth
              multiline
              value={user.password || ""}
              onChange={handleChange}
            />
          </Grid>

          <Grid item xs={4}>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel >คำนำหน้า</InputLabel>
              <Select
                name="title"
                value={user.title || ''}
                onChange={handleChange}
                label="คำนำหน้า"
              >
                {titles.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.titlename}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={8}>
            <TextField
              id="name"
              name="name"
              type="string"
              label="ชื่อ - สกุล"
              variant="outlined"
              fullWidth
              multiline
              value={user.name || ""}
              onChange={handleChange}
            />
          </Grid>

          <Grid item xs={4}>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel >เพศ</InputLabel>
              <Select
                name="sex"
                value={user.sex || ''}
                onChange={handleChange}
                label="เลือกเพศ"
                fullWidth
              >
                {sexs.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.sexname}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={8}>
            <form className={classes.container} noValidate>
              <TextField
                label="เลือกวันเกิด"
                id="birthday"
                name="birthday"
                type="date"
                value={user.birthday || ''} // (undefined || '') = ''
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
                onChange={handleChange}
              />
            </form>
          </Grid>

          <Grid item xs={5}>
            <TextField
              id="age"
              name="age"
              type="string"
              label="กรอกอายุ"
              variant="outlined"
              fullWidth
              multiline
              value={user.age || ""}
              onChange={handleChange}
            />
          </Grid>

          <Grid item xs={7}>
            <TextField
              id="tel"
              name="tel"
              type="string"
              label="กรอกเบอร์โทรศัพท์"
              variant="outlined"
              fullWidth
              multiline
              value={user.tel || ""}
              onChange={handleChange}
            />
          </Grid>

          <Grid item xs={12}>
            <FormControl variant="outlined" className={classes.formControl}>
              <InputLabel >ประเภทสมาชิก</InputLabel>
              <Select
                name="role"
                value={user.role || ''}
                onChange={handleChange}
                label="ประเภทสมาชิก"
                fullWidth
              >
                {roles.map(item => {
                  return (
                    <MenuItem key={item.id} value={item.id}>
                      {item.rolename}
                    </MenuItem>
                  );
                })}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={10}>
            <Button
              name="saveData"
              size="large"
              variant="contained"
              color="primary"
              disableElevation
              className={classes.buttonSty}
              onClick={save}
            >
              บันทึกข้อมูลการสมัครสมาชิก
              </Button>
          </Grid>
        </Grid>
      </Container>
    </div>
  );
};

export default User1;