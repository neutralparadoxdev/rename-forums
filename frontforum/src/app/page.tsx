'use client';
import { FC, useEffect, useState } from "react";
import { PostStubProps } from "./features/PostList/PostStub";
import { PostList } from "./features/PostList/PostList";
import { LoadingComponent } from "./features/loading/Loading";
import { ErrorComponent } from "./features/error/Error";

type HomePageProps = {}

const HomePage: FC<HomePageProps> = () => {
    const [postList, setPostList] = useState<PostStubProps[]>([]);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        setIsLoading(true);
        fetch("/api/group")
        .then(data => data.json())
        .then(data => {
            setIsLoading(false);
            setPostList(data);
        })
        .catch(err => {
            setIsLoading(false);
            console.log(err)
        });
    }, [])

    const page = (
         <main>
            <h1 className="text-4xl">Home</h1>
            <PostList posts={postList} />
        </main>
    )
    return isLoading ? <LoadingComponent />  :
        (error != null ? <ErrorComponent msg={error} /> : page);
}
export default HomePage;