'use client';
import { FC, useState, useEffect } from 'react';
import { LoadingComponent } from '../features/loading/Loading';
import { ErrorComponent } from '../features/error/Error';
import { Header } from '../features/header/header';
import { SignUpLoginModalPurpose } from '../features/signup-login/SignUpLoginModal';

type ForumStubProps = {
    title: string,
    description: string
}

const ForumStub: FC<ForumStubProps>  = ({ title, description }: ForumStubProps) => {
    return (
        <li className="border-2 mt-2 grid grid-cols-[20%_90%]">
            <span className="capitalize text-2xl mr-2 grid-col-1 ">
                <a className="text-black hover:text-red-500" href={"/f/" + title}>{title}</a></span><span className="text mt-auto grid-col-2">{description}</span>
        </li>
    );
}

type ForumsPageProps = {}

const ForumsPage: FC<ForumsPageProps> = () => {
    const [forumList, setForumList] = useState<ForumStubProps[]>([]);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    const [loginSignupState, setLoginSignupState] = useState<SignUpLoginModalPurpose | null>(null);

    useEffect(() => {
        setIsLoading(true);
        fetch("/api/forum")
        .then(data => data.json())
        .then(data => {
            setIsLoading(false);
            setForumList(data)
        })
        .catch(err => {
            setIsLoading(false);
            setError(err)
            console.log("Error loading api: " + err)
        })

    }, [])

    const page = (
        <>
        <Header title="Forums List" link={null} loginSignUpState={loginSignupState} setLoginSignUpState={setLoginSignupState}/>
        <main className="p-2 min-w-full">
            <ol>
                { forumList.toSorted((a, b) => { 
                    if (a.title < b.title) return -1;
                    if (a.title > b.title) return 1;
                    return 0;
                }).map(x => <ForumStub key={x.title} {...x}/>) }
            </ol>
        </main>
        </>
    );

    return (
        isLoading ? <LoadingComponent /> :
        (error !== null ? 
            <ErrorComponent msg={error} /> : page) 
    );
}

export default ForumsPage;