import styles from '@/styles/login.module.css'
import {ChangeEvent, useState} from "react";
import axios from "axios";

type postData = {
    Username  :string
    Password1 :string
    Password2 :string
    Password3 :string
    PinCode   :number
}

export default function Login() {
    const [errorMsg, setErrorMsg] = useState<string>()
    const [userName, setUserName] = useState<string>()
    const [password1, setPassword1] = useState<string>()
    const [password2, setPassword2] = useState<string>()
    const [password3, setPassword3] = useState<string>()
    const [pinCode, setPinCode] = useState<string>()
    const [wait, setWait] = useState(false)

    const changeUserId = (e :ChangeEvent<HTMLInputElement>) => setUserName(e.target.value);
    const changePassword1 = (e :ChangeEvent<HTMLInputElement>) => setPassword1(e.target.value);
    const changePassword2 = (e :ChangeEvent<HTMLInputElement>) => setPassword2(e.target.value);
    const changePassword3 = (e :ChangeEvent<HTMLInputElement>) => setPassword3(e.target.value);
    const changePinCode = (e :ChangeEvent<HTMLInputElement>) => {
        if (e.target.value == null) return;
        if (e.target.value.match( /[^0-9]+/i)) {
            setErrorMsg("pinCode is number")
        } else {
            if (errorMsg == "pinCode is number") setErrorMsg("");
            setPinCode(e.target.value);
        }
    }

    const clickLogin = () => {
        if (userName == undefined || password1 == undefined ||
            password2 == undefined || password3 == undefined || pinCode == undefined) {
            setErrorMsg("userid or passwords or pinCode not entered")
        } else {
            if (errorMsg == "userid or passwords or pinCode not entered") setErrorMsg("");

            const data = {
                Username: userName,
                Password1: password1,
                Password2: password2,
                Password3: password3,
                PinCode: Number(pinCode),
            } as postData

            axios.post("/auth/login", data).then(d => {
                if (d.status == 200) {
                    location.href = "/pr"
                } else {
                    setErrorMsg(d.statusText)
                }
            })
                .catch(e => {
                    setWait(false)
                    setErrorMsg(e.response.data["error"])
                });
            setWait(true)
        }
    }

    return (
        <div className={styles.login}>
            <div className={styles.login_form}>
                {
                    wait ?  (
                        <div className={styles.spinner_base} >
                            <div className={styles.flex}>
                                <div className={styles.spinner}></div>
                            </div>
                            <div className={styles.flex}>
                                <span className={styles.spinner_text}>Please Wait...</span>
                            </div>
                        </div>
                    ) : null
                }
                <div className="login_form_top">
                    <h1>LOGIN</h1>
                    <p>{errorMsg}</p>
                </div>
                <div className="login_form_btm">
                    <input type="id" name="user_id" placeholder="UserID" onChange={changeUserId} required={true} />
                    <input type="password" name="password" placeholder="Password1" onChange={changePassword1} required={true}/>
                    <input type="password" name="password" placeholder="Password2" onChange={changePassword2} required={true}/>
                    <input type="password" name="password" placeholder="Password3" onChange={changePassword3} required={true}/>
                    <input type="password" name="password" placeholder="PinCode" onChange={changePinCode} required={true} />
                    <button type="button" name="botton" onClick={clickLogin} disabled={wait}>LOGIN</button>
                </div>
            </div>
        </div>
    )
}