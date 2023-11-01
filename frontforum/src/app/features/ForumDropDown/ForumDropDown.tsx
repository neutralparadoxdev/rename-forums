'use client';

import { usePathname } from "next/navigation";
import { useRouter } from "next/navigation";
import { ChangeEvent, ChangeEventHandler, FC, FormEvent, useEffect, useState } from "react";

function getForumFromPath(path: string): string {
    if(path === "/") {
        return "";
    }
    const list =  path.split("/")

    return list[list.length-1]
}

export const ForumDropDown: FC<{}> = () => {
    const pathname = usePathname();
    const router = useRouter();
    const forumName = getForumFromPath(pathname);

    const [forumSelected, setForumSelect] = useState<string>("");

    useEffect(() => {
        console.log(`path: ${pathname}\nname: ${forumName}`)
    })

    function submit(e: FormEvent<HTMLFormElement>) {
        e.preventDefault();
        if(forumSelected === "home") {
            router.push("/")
            return
        }
        router.push("/f/" + forumSelected);
        return
    }

    const forumSelecteFieldOnChange: ChangeEventHandler<HTMLInputElement> = (e: ChangeEvent<HTMLInputElement>) => {
        if(/\s/.test(e.target.value)) {
            return
        }
        e.preventDefault();
        setForumSelect(e.target.value.toLocaleLowerCase());
    }

    return (
        <form onSubmit={submit}>
        <input 
            onChange={forumSelecteFieldOnChange} 
            value={forumSelected} 
            list="selection-forum-list" 
            className="border-2 w-32 ml-2 p-2 text-center" 
            placeholder={forumName === "" ? "home" : forumName} /> 
        <datalist id="selection-forum-list">
            <option value="home"></option>
            <option value="science"></option>
            <option value="math"></option>
        </datalist>
        </form>
    );
}
