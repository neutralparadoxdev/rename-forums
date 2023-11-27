'use client';

import { ErrorComponent } from "@/app/features/error/Error";
import { Header } from "@/app/features/header/header";
import { LoadingComponent } from "@/app/features/loading/Loading";
import { SignUpLoginModalPurpose } from "@/app/features/signup-login/SignUpLoginModal";
import { usePathname } from "next/navigation";
import { FC, FormEvent, useEffect, useState } from "react";
import { useRouter } from 'next/navigation';
import { GetSessionToken, SessionTokenExist } from "@/app/services/SessionManager/session";
import Markdown from "react-markdown";

type PostPageProps = {}


type PostResponse = {
    title: string,
    body: string,
    authorName: string,
}

function getForumPathParam(pathname : string) {
  const pathSplit = pathname.split('/')
  const postParam = pathSplit[pathSplit.length-1]
  const forumParam = pathSplit[pathSplit.length-2]
  return [postParam, forumParam]
}

type PostEditor = {
    title: string,
    body: string,
    id: string,
    author: string,
    form: string,
    update: (title: string, body: string) => void,
    close: () => void
}

const PostEditor: FC<PostEditor> = ({title, body, id, form, update, close} : PostEditor) => {
    const [editTitle, setEditTitle] = useState<string>(title);
    const [editBody, setEditBody] = useState<string>(body);

    const router = useRouter();

    async function submit(e: FormEvent<HTMLFormElement>) {
        e.preventDefault();


        if (!SessionTokenExist()) {
            console.log("submit:no session")
            return
        }

        const sessionToken = GetSessionToken()!

        if(editTitle.length === 0) {
            // cant be submitted
            return;
        }

        console.log(`/api/post/${form}/${id}`)
        console.log(JSON.stringify({
                title: title === editTitle ? null : editTitle,
                body: body === editBody ? null : editBody
            }))
        try {
            const req = await fetch(`/api/post/${form}/${id}`, {
                method: "PATCH",
                headers: {
                    'Content-Type': "application/json",
                    'Bearer-Token': sessionToken
                },
                body: JSON.stringify({
                    title: title === editTitle ? null : editTitle,
                    body: body === editBody ? null : editBody
                })
            })
            if(req.status === 204) {
                router.refresh()
                update(editTitle,editBody)
                close()
            } else {
                console.log(req.status)
            }
        }
        catch(err) {
            console.log(err);
        }
    }

    return (
    <form className="border-2 border-purple-700 p-2 flex flex-col" onSubmit={submit}>
        <h2 className="text-2xl font-bold text-purple-700">Editing Post</h2>
        <label className="text-xl font-bold text-purple-700" htmlFor="title-edit">Title</label>
        <input 
            id="title-edit"
            type="text" 
            className="border-2 p-2 border-purple-700 text-2xl font-bold mb-2" 
            onChange={(e) => setEditTitle(e.target.value)}
            value={editTitle} />
        <label className="text-xl font-bold text-purple-700" htmlFor="body-edit">Body</label>
        <textarea 
            id="body-edit"
            className="border-2 p-4 border-2 border-purple-700" 
            onChange={(e) => setEditBody(e.target.value)}
            value={editBody} />
        <div className="ml-auto mr-2">
            <button 
                onClick={()  => close()}
                className="mr-2 border-2 p-2 mt-2 w-32 border-purple-700 text-purple-700 hover:font-bold">Cancel</button>
            <button 
                className="border-2 p-2 mt-2 w-32 border-purple-700 text-purple-700 hover:font-bold">Save</button>
        </div>
    </form>

    );
}

const PostPage: FC<PostPageProps> = () => {
    const pathname = usePathname();
    const [postId, forumName] = getForumPathParam(pathname);
    const [post, setPost] = useState<PostResponse | null>(null);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);
    const [loginSignUpState, setLoginSignUpState] = useState<SignUpLoginModalPurpose | null>(null);

    const [isEditing, setIsEditing] = useState<boolean>(false);

    const [refresh, setRefresh] = useState<boolean>(false);

    const [username, setUsername] = useState<string | null>(null);

    const sessionToken = GetSessionToken();

    const [deletePostPrompt, setDeletePostPrompt] = useState<boolean>(false);

    const router = useRouter();

    function setPromptDeletePost() {
        setDeletePostPrompt(true);
    }

    useEffect(() => {
        
        setIsLoading(true);
        if(sessionToken === null || sessionToken === "") {

            fetch("/api/post/" + forumName + "/" + postId)
            .then(data => data.json())
            .then(data => {
                setIsLoading(false);
                setPost(data);
            })
            .catch(err => {
                setIsLoading(false);
                setError("Not Found");
                console.log(err)
            })
        } else {
            fetch("/api/post/" + forumName + "/" + postId, {
                headers: {
                    "Bearer-Token": sessionToken
                }
            })
            .then(data => data.json())
            .then(data => {
                setIsLoading(false);
                setPost(data);
            })
            .catch(err => {
                setIsLoading(false);
                setError(err);
                setError("Page Not Found")
                console.log(err)
            })

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
                setError("Page Not Found")
                console.log(err);
            })
        }
    }, [refresh])

    async function deletePost() {

        if(sessionToken === null || sessionToken === "") {
            console.log("TRIED TO DELETE WITHOUT TOKEN")
            return
        }

        let res = await fetch(`/api/post/${forumName}/${postId}`, {
            method: "DELETE",
            headers: {
                "Bearer-Token" : sessionToken
            }
        })

        if (res.status !== 204) {
            console.log("COULD NOT DELETE")
        } else {
            router.push(`/f/${forumName}`)
        }
    }

    const page = post !== null ? (
        <>
            {
                deletePostPrompt ? 
                <div className="w-screen h-screen bg-red-400 absolute">
                    <div className="w-min-full h-min-full flex flex-col items-center justify-center mt-40">
                        <div className="bg-white max-w-fit m-4 p-8">
                            <h1 className="text-2xl text-center mt-4">Do you want to delete the Post?</h1>
                            <h3 className="text-xl text-center mt-4">This is a permanent and not reversible</h3>
                            <div className="flex flex-row justify-between mt-20">
                                <button className="p-4 border-2 border-red-400 w-24 text-red-400"
                                onClick={async () => deletePost()}>Delete</button>
                                <button className="p-4 border-2 text-gray-400 w-24"
                                    onClick={() => setDeletePostPrompt(false)}>Cancel</button>
                            </div>
                        </div>
                    </div>
                </div> :
                <></>
            }

            <Header title={forumName} link={"/f/" + forumName} loginSignUpState={loginSignUpState} setLoginSignUpState={(x) => { setLoginSignUpState(x); }}/>
            <main className="m-2">
                { isEditing ? 
                    <PostEditor 
                        form={forumName} 
                        author={post.authorName} 
                        title={post.title} 
                        body={post.body} 
                        id={postId} 
                        update={(title, body) => setPost(val => { return { title: title, body: body, authorName: val!.authorName } })}
                        close={() => setIsEditing(false)}
                        /> :
                <>
                <div className="flex">
                    <div className="flex flex-col justify-center">
                        <h2 className="text-2xl font-bold">{post.title}</h2>
                    </div>
                    { sessionToken !== "null" && sessionToken !== "" && username !== null && username === post.authorName ?
                    <>
                    <button 
                        className="border-2 p-2 m-2 w-24 hover:text-blue-400 hover:border-blue-400 hover:font-bold" 
                        onClick={() => setIsEditing(true)}>Edit</button>
                    <button 
                        className="border-2 p-2 m-2 border-red-400 text-red-400 hover:border-red-600 hover:text-red-600 hover:font-bold"
                        onClick={() => setPromptDeletePost()}
                        >Delete</button>
                    </>
                        : <></>
                    }
                </div>
                <h3 className="text-blue-400">By <a className="text-blue-400 hover:text-red-400" href="/">{post.authorName}</a></h3>
                <Markdown className="border-2 p-4">{post.body}</Markdown>
                </> }
                <div className="mt-4 border-2 p-4">
                    <form className="min-w-full">
                        <label className="block" htmlFor="comment-field" >Please Leave a comment</label>
                        <input id="comment-field" className="border-yellow-500 border-2 w-4/5 ml-auto mr-auto p-1" type="text" />
                    </form>
                    <h4>Comments Go Here</h4>
                </div>
            </main>
        </>
    ) : <></>;
    return (isLoading ? <LoadingComponent /> :
        (error !== null ? <ErrorComponent msg={error} /> : page));

}

export default PostPage;
