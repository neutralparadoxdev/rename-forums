'use client';

import { FC, useState } from "react"
import { Header } from "../header/header";

import { SignUpLoginModalPurpose } from "../signup-login/SignUpLoginModal";

export type ErrorProps = {
    msg: string
}

export const ErrorComponent: FC<ErrorProps> = ({ msg } : ErrorProps) => {
    const [loginSignUpState, setLoginSignUpState] = useState<SignUpLoginModalPurpose | null>(null);
    
    return (
        <main className="w-min-max">

            <Header title="Error" link={null} loginSignUpState={loginSignUpState} setLoginSignUpState={setLoginSignUpState}/>
            <h1 className="text-4xl w-min-max text-center pt-20">Error</h1>
            <h2 className="text-2xl w-min-max text-center pt-20">{msg}</h2>
        </main>
    );
}