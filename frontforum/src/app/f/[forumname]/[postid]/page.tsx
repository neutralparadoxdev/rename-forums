'use client';

import { ErrorComponent } from "@/app/features/error/Error";
import { Header } from "@/app/features/header/header";
import { LoadingComponent } from "@/app/features/loading/Loading";
import { SignUpLoginModalPurpose } from "@/app/features/signup-login/SignUpLoginModal";
import { usePathname } from "next/navigation";
import { ChangeEvent, FC, FormEvent, useEffect, useState } from "react";
import { useRouter } from 'next/navigation';

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

const PostEditor: FC<PostEditor> = ({title, body, id, author, form, update, close} : PostEditor) => {
    const [editTitle, setEditTitle] = useState<string>(title);
    const [editBody, setEditBody] = useState<string>(body);

    const router = useRouter();

    async function submit(e: FormEvent<HTMLFormElement>) {
        e.preventDefault();

        const sessionToken = localStorage.getItem('session-token')

        if (sessionToken === null || sessionToken === "") {
            console.log("submit:no session")
            return
        }

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
            className="border-2 p-2 border-2 border-purple-700" 
            onChange={(e) => setEditBody(e.target.value)}
            value={editBody} />
        <button 
            className="ml-auto mr-2 border-2 p-2 mt-2 w-32 border-purple-700 text-purple-700 hover:font-bold">Save</button>
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

    const sessionToken = localStorage.getItem('session-token')

    function setPromptDeletePost() {
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
                setError(err);
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
                console.log(data)
                setIsLoading(false);
                setPost(data);
            })
            .catch(err => {
                setIsLoading(false);
                setError(err);
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
                console.log(err);
            })
        }
    }, [refresh])

    const page = post !== null ? (
        <>
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
                <h3 className="text-blue-400">By <a className="hover:text-red-400" href="/">{post.authorName}</a></h3>
                <p className="border-2 p-4">{post.body}</p>
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