'use client';

import { Header } from "@/app/features/header/header";
import { usePathname } from "next/navigation";
import { FC } from "react";

function getUserNameFromPath(pathname : string) {
  const pathSplit = pathname.split('/')
  const username = pathSplit[pathSplit.length-1]
  return username;
}


const UserPage: FC<{}> = () => {
    const pathname = usePathname();
    const username = getUserNameFromPath(pathname)
    
    return (
    <>
        <Header title={username} link={null} />
        <main>
            <h3>Created at: Sometime</h3>
            <h3>Last Logged in: Sometime</h3> {/* can be hidden */}
        </main>
    </>);
}

export default UserPage;