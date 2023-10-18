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
         <main>
            <h1>Post</h1>
            <h2>{post.title}</h2>
            <h3>By {post.authorName}</h3>
            <p>{post.body}</p>
        </main>
    ) : <></>;
    return (isLoading ? <LoadingComponent /> :
        (error !== null ? <ErrorComponent msg={error} /> : page));

}

export default PostPage;