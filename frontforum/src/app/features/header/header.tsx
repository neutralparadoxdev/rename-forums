import { useSearchParams } from "next/navigation";
import { FC, useState, useEffect } from "react";
import { SignUpLoginModal, SignUpLoginModalPurpose } from "../signup-login/SignUpLoginModal";
import { useRouter } from "next/navigation";
import { ForumDropDown } from '../ForumDropDown/ForumDropDown';
import Link from "next/link";


export type HeaderProps = {
    title: string,
    link: string | null,
    loginSignUpState: SignUpLoginModalPurpose | null, 
    setLoginSignUpState: (purpose: SignUpLoginModalPurpose | null) => void,
    sessionRequired?: boolean
}

export const Header: FC<HeaderProps> = ({title, link, loginSignUpState, setLoginSignUpState, sessionRequired} : HeaderProps) => {
    const [username, setUsername] = useState<string>("");
    const router = useRouter();

    const sessionToken = localStorage.getItem('session-token')

    const [reload, triggerReload] = useState<boolean>(false);

    useEffect(() => {
        if(sessionRequired !== undefined && (sessionToken === null || sessionToken === "") && loginSignUpState === null) {
            setLoginSignUpState(SignUpLoginModalPurpose.Login)
        }

    }, [loginSignUpState])

    function logout() {
        const token = localStorage.getItem('session-token')

        if(token !== null) {
            fetch('/api/session',{
                method: 'DELETE',
                headers: {
                    'Bearer-Token' : token,
                },
            })
            .finally(() => {
                localStorage.setItem('session-token', "")
                router.refresh()
                triggerReload(val => !val);
            })
        }
        if(sessionRequired !== undefined) {
            setLoginSignUpState(SignUpLoginModalPurpose.Login);
        }
    }

    useEffect(() => {
        if(sessionToken !== null &&  sessionToken != "") {
            fetch('/api/me', {
                method: "GET",
                headers: {
                    'Bearer-Token': sessionToken
                },
            })
            .then(data => data.json())
            .then(data => {
                setUsername(data.username);
            })
            .catch(err => {
                console.log(err);
            })
        }
    }, [sessionToken])

    function TriggerLogin() {
        setLoginSignUpState(SignUpLoginModalPurpose.Login);
    }

    return (<>
        { loginSignUpState != null ? 
        <SignUpLoginModal 
            purpose={loginSignUpState} 
            close={() => {setLoginSignUpState(null);}} 
            changePurpose={(purpose) => setLoginSignUpState(purpose)}
            setAuthToken={(token) => { setLoginSignUpState(null); }}
            /> : 
            <></> }
    <header className="border-b-4 border-[blue] flex justify-between p-3">
        <div className="flex flex-row">
        { link !== null ? <h1 className="capitalize text-3xl font-bold"><a href={link}> {title}</a></h1> 
        : <h1 className="capitalize text-3xl font-bold">{title}</h1> }
        <ForumDropDown />
        </div>

        <div className="flex flex-row min-w-fit justify-between gap-1 align-center">
            { sessionToken === null || sessionToken === "" ?
            <>
            <button className="border-2 text-left min-w-fit p-1" onClick={TriggerLogin}>Create Forum</button>
            <button className="border-2 hover:text-red-400 min-w-fit p-1" onClick={() => setLoginSignUpState(SignUpLoginModalPurpose.SignUp)}>Sign Up</button>
            <button className="border-2 hover:text-red-400 min-w-fit p-1" onClick={() => setLoginSignUpState(SignUpLoginModalPurpose.Login)}>Log In</button>
            </> : 
            <>
            <div className="border-2 flex flex-col justify-center">
                <Link className="text-center align-middle content-center" href="/forum/new">Create Forum</Link>
            </div>
            <div className="border-2 align-middle flex flex-col justify-center"><p className="max-h-fit max-w-fit">{ `User: ${username}` }</p></div>
            <button className="border-2 p-1 hover:text-red-400" onClick={logout}>Logout</button>
            </>
            }
        </div>
    </header>
    </>)
}