import { useSearchParams } from "next/navigation";
import { FC, useState } from "react";
import { SignUpLoginModal, SignUpLoginModalPurpose } from "../signup-login/SignUpLoginModal";

export type HeaderProps = {
    title: string,
    link: string | null,
}

export const Header: FC<HeaderProps> = ({title, link} : HeaderProps) => {

    const [loginSignupPrompt, setLoginSignupPrompt] = useState<SignUpLoginModalPurpose | null>();

    return (<>
        { loginSignupPrompt != null ? 
        <SignUpLoginModal 
            purpose={loginSignupPrompt} 
            close={() => {setLoginSignupPrompt(null);}} 
            changePurpose={(purpose) => setLoginSignupPrompt(purpose)}
            /> : 
            <></> }
    <header className="border-b-4 border-[blue]  flex justify-between pr-2 pl-2">
        { link !== null ? <h1 className="capitalize text-3xl font-bold"><a href={link}> {title}</a></h1> 
        : <h1 className="capitalize text-3xl font-bold">{title}</h1> }

        <div className="mt-auto mb-0">
            <button className="mr-2 hover:text-red-400" onClick={() => setLoginSignupPrompt(SignUpLoginModalPurpose.SignUp)}>Sign Up</button>
            <button className="hover:text-red-400" onClick={() => setLoginSignupPrompt(SignUpLoginModalPurpose.Login)}>Log In</button>
        </div>
    </header>
    </>)
}