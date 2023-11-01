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
    setLoginSignUpState: (purpose: SignUpLoginModalPurpose | null) => void
}

export const Header: FC<HeaderProps> = ({title, link, loginSignUpState, setLoginSignUpState} : HeaderProps) => {
    const [username, setUsername] = useState<string>("");
    const router = useRouter();

    const sessionToken = localStorage.getItem('session-token')

    const [reload, triggerReload] = useState<boolean>(false);

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

    return (<>
        { loginSignUpState != null ? 
        <SignUpLoginModal 
            purpose={loginSignUpState} 
            close={() => {setLoginSignUpState(null);}} 
            changePurpose={(purpose) => setLoginSignUpState(purpose)}
            setAuthToken={(token) => { setLoginSignUpState(null); }}
            /> : 
            <></> }
    <header className="border-b-4 border-[blue]  flex justify-between p-3">
        <div className="flex flex-row">
        { link !== null ? <h1 className="capitalize text-3xl font-bold"><a href={link}> {title}</a></h1> 
        : <h1 className="capitalize text-3xl font-bold">{title}</h1> }
        <ForumDropDown />
        </div>

        <div className="mt-auto mb-0">
            { sessionToken === null || sessionToken === "" ?
            <>
            <Link className="border-2 mb-6 p-1 w-24" href="/forum/new">Create a Forum</Link>
            <button className="mr-2 hover:text-red-400" onClick={() => setLoginSignUpState(SignUpLoginModalPurpose.SignUp)}>Sign Up</button>
            <button className="hover:text-red-400" onClick={() => setLoginSignUpState(SignUpLoginModalPurpose.Login)}>Log In</button>
            </> : 
            <>
            <Link className="border-2 mb-6 p-1 w-24" href="/forum/new">Create a Forum</Link>
            <span className="mr-2">{ username }</span>
            <button className="mr-2 hover:text-red-400" onClick={logout}>Logout</button>
            </>
            }
        </div>
    </header>
    </>)
}