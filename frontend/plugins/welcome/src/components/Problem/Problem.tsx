import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';

import { EntProblem } from '../../api/models/EntProblem';
import { EntProblemStatus } from '../../api/models/EntProblemStatus';
import { EntProblemTitle } from '../../api/models/EntProblemTitle';
import { EntRoom } from '../../api/models/EntRoom';
import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
  }),
);

export default function Create() {
  const classes = useStyles();
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบเเจ้งปัญหา' };
  const api = new DefaultApi();

  //const [deceased, setDeceased] = useState(String);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [user, setUser] = useState<EntUser[]>([]);
  const [room, setRoom] = useState<EntRoom[]>([]);
  const [problemtitle, setProblemtitle] = useState<EntProblemTitle[]>([]);
  const [problemstatus, setProblemstatus] = useState<EntProblemStatus[]>([]);
  const [loading, setLoading] = useState(true);
  const [datetime, setDatetime] = useState(String);

  const [userid, setUsers] = useState(Number);
  const [roomid, setRooms] = useState(Number);
  const [problemtitleid, setProblemtitles] = useState(Number);
  const [probleminfo, setProbleminfos] = useState(String);
  const [problemstatusid, setProblemstatuss] = useState(Number);

  useEffect(() => {

    const getUserid= async () => {
      const res = await api.listUser({ limit: 3, offset: 0 });
      setLoading(false);
      setUser(res);
    };
    getUserid();

    const getRoomid = async () => {
      const res = await api.listRoom({ limit: 2, offset: 0 });
      setLoading(false);
      setRoom(res);
    };
    getRoomid();

    const getProblemtitle = async () => {
      const res = await api.listProblemtitle({ limit: 6, offset: 0 });
      setLoading(false);
      setProblemtitle(res);
    };
    getProblemtitle();

    const getProblemstatus = async () => {
      const res = await api.listProblemstatus({ limit: 4, offset: 0 });
      setLoading(false);
      setProblemstatus(res);
    };
    getProblemstatus();

  }, [loading]);
  

  const userhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUsers(event.target.value as number);
  };

  const datetimehandleChange = (event: any) => {
    setDatetime(event.target.value as string);
  };

  const roomidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setRooms(event.target.value as number);
  };

  const problemtitleidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProblemtitles(event.target.value as number);
  };

  const ProbleminfohandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProbleminfos(event.target.value as string);
  };
  const ProblemstatusidhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setProblemstatuss(event.target.value as number);
  };

  const CreateProblem = async () => {
    const problem = {
      user:userid,
      added: datetime + ":00+07:00" , 
      probleminfo: probleminfo,
      problemstatus: problemstatusid,
      problemtitle: problemtitleid,
      room: roomid, 
    };
    console.log(problem);
    const res: any = await api.createProblem({ problem : problem });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      window.location.reload(false);
    } else {
      setAlert(false);
    }
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      //subtitle="Some quick intro and links."
      ></Header>
      <Content>
        <ContentHeader title="">
          <Typography align="left" style={{ marginRight: 16, color: "#00eeff" }}>
          </Typography>
          <div>
            <Button variant="contained" color="primary">
              ออกจากระบบ
         </Button>
          </div>
          
        </ContentHeader>
        <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel>สมาชิกระบบหอพัก</InputLabel>
                <Select
                  id="user"
                  value={userid}
                  onChange={userhandleChange}
                  style={{ width: 400 }}
                >
                {user.map((item:EntUser ) => (
                  <MenuItem value={item.id}>{item.id}</MenuItem>
                ))}
                </Select>
              </FormControl>
            </div>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel>เลขห้องพัก</InputLabel>
                <Select
                  id="room"
                  value={roomid}
                  onChange={roomidhandleChange}
                  style={{ width: 400 }}
                >
                {room.map((item: EntRoom) => (
                  <MenuItem value={item.id}>{item.id}</MenuItem>
                ))}
                </Select>
              </FormControl>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel>หัวข้อปัญหา</InputLabel>
                <Select
                  id="problemtitle"
                  value={problemtitleid}
                  onChange={problemtitleidhandleChange}
                  style={{ width: 400 }}
                >
                {problemtitle.map((item: EntProblemTitle) => (
                  <MenuItem value={item.id}>{item.problemtitle}</MenuItem>
                ))}
                </Select>
              </FormControl>
            </div>

            <FormControl>
            <TextField
               name="probleminfo"
               label="รายละเอียดของปัญหา"
               variant="outlined"
               type="string"
               size="medium"
               value={probleminfo}
               onChange={ProbleminfohandleChange}
             />
             </FormControl>

             <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <InputLabel>สถานะของคุณ</InputLabel>
                <Select
                  id="problemstatus"
                  value={problemstatusid}
                  onChange={ProblemstatusidhandleChange}
                  style={{ width: 400 }}
                >
                {problemstatus.map((item: EntProblemStatus) => (
                  <MenuItem value={item.id}>{item.problemstatus}</MenuItem>
                ))}
                </Select>
              </FormControl>
            </div>
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <TextField
              id="datetime"
              label="ว/ด/ป "
              type="datetime-local"
              value={datetime}
              onChange={datetimehandleChange}
              InputLabelProps={{
                shrink: true,
              }}
            />
              </FormControl>
            </div>

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateProblem();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยัน
             </Button>
             {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  บันทึกสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    พบปัญหาระหว่างบันทึกข้อมูล
                  </Alert>
                )}
            </div>
          ) : null}
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                กลับ
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}