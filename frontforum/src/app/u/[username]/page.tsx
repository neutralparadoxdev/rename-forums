'use client';

import { Header } from "@/app/features/header/header";
import { usePathname } from "next/navigation";
import { useState, FC } from "react";
import {SignUpLoginModalPurpose} from '../../features/signup-login/SignUpLoginModal';

function getUserNameFromPath(pathname : string) {
  const pathSplit = pathname.split('/')
  const username = pathSplit[pathSplit.length-1]
  return username;
}

const UserPage: FC<{}> = () => {
    const pathname = usePathname();
    const username = getUserNameFromPath(pathname)
    const [loginSignUpState, setLoginSignUpState] = useState<SignUpLoginModalPurpose | null>(null);
   
    return (
    <>
        <Header 
			title={username} 
			link={null} 
			loginSignUpState={loginSignUpState}
			setLoginSignUpState={setLoginSignUpState}
		/>
        <main>
            <h3>Created at: Sometime</h3>
            <h3>Last Logged in: Sometime</h3> 
        </main>
    </>);
}

export default UserPage;
