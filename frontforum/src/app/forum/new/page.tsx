'use client';

import { FC, useState } from 'react';

import { Header } from '@/app/features/header/header';
import { SignUpLoginModalPurpose } from '@/app/features/signup-login/SignUpLoginModal';

enum PublicityState {
    PUBLIC,
    PRIVATE,
}

const NewForumPage: FC<{}> = () => {
    const [loginSignupState, setLoginSignUpState] = useState<SignUpLoginModalPurpose | null>(null);

    const [title, setTitle] = useState<string>("");
    const [description, setDescription] = useState<string>("");
    const [publicity, setPublicity] = useState<PublicityState>(PublicityState.PRIVATE)

    const [titleMessage, setTitleMessage] = useState<boolean>(false);
    
    function processTitleChange(text: string) {
        setTitle(text)
    }

    return ( 
    <>
    <Header  link={null} title="Create a Forum" loginSignUpState={loginSignupState} setLoginSignUpState={setLoginSignUpState}/>
    <main>
        <form className="p-2">
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