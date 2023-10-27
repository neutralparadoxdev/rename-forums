import { useSearchParams } from "next/navigation";
import { FC, useState } from "react";
import { SignUpLoginModal, SignUpLoginModalPurpose } from "../signup-login/SignUpLoginModal";
import { useRouter } from "next/navigation";

export type HeaderProps = {
    title: string,
    link: string | null,
}

export const Header: FC<HeaderProps> = ({title, link} : HeaderProps) => {
    const router = useRouter();

    const sessionToken = localStorage.getItem('session-token')

    const [loginSignupPrompt, setLoginSignupPrompt] = useState<SignUpLoginModalPurpose | null>();

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
            })
        }
    }

    return (<>
        { loginSignupPrompt != null ? 
        <SignUpLoginModal 
            purpose={loginSignupPrompt} 
            close={() => {setLoginSignupPrompt(null);}} 
            changePurpose={(purpose) => setLoginSignupPrompt(purpose)}
            setAuthToken={(token) => { setLoginSignupPrompt(null); }}
            /> : 
            <></> }
    <header className="border-b-4 border-[blue]  flex justify-between pr-2 pl-2">
        { link !== null ? <h1 className="capitalize text-3xl font-bold"><a href={link}> {title}</a></h1> 
        : <h1 className="capitalize text-3xl font-bold">{title}</h1> }

        <div className="mt-auto mb-0">
            { sessionToken === null || sessionToken === "" ?
            <>
            <button className="mr-2 hover:text-red-400" onClick={() => setLoginSignupPrompt(SignUpLoginModalPurpose.SignUp)}>Sign Up</button>
            <button className="hover:text-red-400" onClick={() => setLoginSignupPrompt(SignUpLoginModalPurpose.Login)}>Log In</button>
            </> : 
            <>
            <button className="mr-2 hover:text-red-400" onClick={logout}>Logout</button>
            </>
            }
        </div>
    </header>
    </>)
}