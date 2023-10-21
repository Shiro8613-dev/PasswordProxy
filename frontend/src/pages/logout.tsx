import styles from '@/styles/logout.module.css';
import axios from "axios";
import {useEffect, useState} from "react";

export default function Logout() {
    const [after, setAfter] = useState(false);
    const noClick = () => location.href = "/";
    const yesClick = () => {
        new Promise<void>( async (resolve, reject) => {
            const res = await axios.post("/auth/logout");
            if (res.status != 200) {
                reject();
            } else {
                resolve();
            }
        }).then(() => setAfter(true));
    }

    useEffect(() => {
        if (after) {
            const s = setTimeout(() => {
                location.href = "/auth/login"
                clearTimeout(s)
            },5 * 1000)
        }
    }, [after])

    return (
        <div className={styles.logout}>
            {
                after ? (
                    <div className={styles.box}>
                        <div className={`${styles.flex} ${styles.just}`}>
                            <span className={styles.text}>ログアウトしました。</span>
                        </div>
                    </div>
                ):(
                    <div className={styles.box}>
                        <div className={`${styles.flex} ${styles.just}`}>
                            <span className={styles.text}>ログアウトを実行しますか？</span>
                        </div>
                        <div className={`${styles.flex} ${styles.just}`}>
                            <div className={styles.buttons}>
                                <div className={styles.button} onClick={noClick}>いいえ</div>
                                <div className={styles.button} onClick={yesClick}>はい</div>
                            </div>
                        </div>
                    </div>
                )
            }
        </div>
    )
}