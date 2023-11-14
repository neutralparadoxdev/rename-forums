'use client';

import { FC, FormEvent, useState } from 'react';

import { Header } from '@/app/features/header/header';
import { SignUpLoginModalPurpose } from '@/app/features/signup-login/SignUpLoginModal';
import { useRouter } from 'next/navigation';
import { GetSessionToken } from '@/app/services/SessionManager/session';

enum PublicityState {
    PUBLIC,
    PRIVATE,
}

const NewForumPage: FC<{}> = () => {
    const sessionToken = GetSessionToken();
    const [loginSignupState, setLoginSignUpState] = useState<SignUpLoginModalPurpose | null>(
        (sessionToken === null || sessionToken === "") ? SignUpLoginModalPurpose.Login : null)

    const [title, setTitle] = useState<string>("");
    const [description, setDescription] = useState<string>("");
    const [publicity, setPublicity] = useState<PublicityState>(PublicityState.PRIVATE)

    const [titleMessage, setTitleMessage] = useState<boolean>(false);

    const router = useRouter();
    
    function processTitleChange(text: string) {
        if(/\s/.test(text)) {
            return
        }
        setTitle(text)
    }

    async function submit(e: FormEvent<HTMLFormElement>) {
        e.preventDefault();

        if(sessionToken === null || sessionToken === "") {
            return
        }

        try {
            const res = await fetch("/api/forum", {
                method: "POST",
                headers: {
                    "Bearer-Token" : sessionToken,
                    "Content-Type" : "application/json"
                },
                body: JSON.stringify({
                    "title": title,
                    "description": description,
                    "is_public" : publicity === PublicityState.PRIVATE ? false : true,
                })
            })

            if (res.status === 204) {
                router.push(`/f/${title}`)
                return;
            }
        }
        catch(err) {
            console.log(err)
        }


    }

    return ( 
    <>
    <Header 
        link={null} 
        title="Create a Forum" 
        loginSignUpState={loginSignupState} 
        setLoginSignUpState={setLoginSignUpState}
        sessionRequired={true}
    />
    <main>
        <form className="p-2" onSubmit={submit}>
            <label htmlFor="title">Title</label>
            <input value={title} onChange={x => processTitleChange(x.target.value)} id="title" name="title" type="text" className="border-2 p-2 mb-2 min-w-full" required/>
            <label htmlFor="description">Description</label>
            <textarea value={description} onChange={x => setDescription(x.target.value)} id="description" name="description" rows={10} className="border-2 p-2 mb-2 min-w-full" required/>
            <fieldset className="border-2 m-2 p-2 max-w-fit">
                <legend>Publicity Status:</legend>
                <label className="block" htmlFor="public">
                    <input 
                        type="radio" 
                        name="openness" 
                        id="public" 
                        value="public"
                        checked={publicity == PublicityState.PUBLIC}
                        onChange={() => setPublicity(PublicityState.PUBLIC)}
                        />
                    Public
                </label>
                <label htmlFor="private">
                    <input 
                        type="radio" 
                        name="openness" 
                        id="private" 
                        value="private" 
                        checked={publicity == PublicityState.PRIVATE}
                        onChange={() => setPublicity(PublicityState.PRIVATE)}/>
                    Private
                </label>
            </fieldset>
            <button className="border-2 w-24 m-2 p-2 hover:text-red-400 hover:border-red-400">Submit</button>
        </form>
    </main>
    </>
    );
}

export default NewForumPage;