'use client';

import { ErrorComponent } from "@/app/features/error/Error";
import { LoadingComponent } from "@/app/features/loading/Loading";
import { usePathname } from "next/navigation";
import { FC, useEffect, useState } from "react";

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

const PostPage: FC<PostPageProps> = () => {
    const pathname = usePathname();
    const [postId, forumName] = getForumPathParam(pathname);
    const [post, setPost] = useState<PostResponse | null>(null);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        setIsLoading(true);
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
    }, [])

    const page = post !== null ? (
        <>
            <header className="border-b-4 border-[blue]  flex justify-between pr-2 pl-2">
                <h1 className="capitalize text-3xl font-bold">{forumName}</h1>
                <div className="mt-auto mb-0">
                    <a className="mr-2 hover:text-red-400" href="">Sign In</a>
                    <a className="hover:text-red-400" href="">Log In</a>
                </div>
            </header>
            <main className="m-2">
                <h2 className="text-2xl font-bold">{post.title}</h2>
                <h3 className="text-blue-400">By <a className="hover:text-red-400" href="/">{post.authorName}</a></h3>
                <p className="border-2 p-4">{post.body}</p>
                <div className="mt-4 border-2 p-4">
                    <form className="min-w-full">
                        <label htmlFor="comment-field" >Please Leave a comment</label>
                        <input id="comment-field" className="border-2 w-4/5 ml-auto mr-auto p-1" type="text" />
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