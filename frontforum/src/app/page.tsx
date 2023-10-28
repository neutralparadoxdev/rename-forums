'use client';
import { FC, useEffect, useState } from "react";
import { PostStubProps } from "./features/PostList/PostStub";
import { PostList } from "./features/PostList/PostList";
import { LoadingComponent } from "./features/loading/Loading";
import { ErrorComponent } from "./features/error/Error";
import { Header } from "./features/header/header";
import { NewPost } from "./features/NewPost/NewPost";
import { SignUpLoginModalPurpose } from "./features/signup-login/SignUpLoginModal";

type HomePageProps = {}

const HomePage: FC<HomePageProps> = () => {
    const [postList, setPostList] = useState<PostStubProps[]>([]);
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    const [loginSignUpState, setLoginSignUpState] = useState<SignUpLoginModalPurpose | null>(null);

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
        <>
        <Header title="Home" link={null} setLoginSignUpState={(x) => { setLoginSignUpState(x);}} loginSignUpState={loginSignUpState} />
        <main>
            <NewPost rows={8} forums={["math", "science"]} showLogin={() => {setLoginSignUpState(SignUpLoginModalPurpose.Login)}}/>
            <PostList posts={postList} />
        </main>
        </>
    )
    return isLoading ? <LoadingComponent />  :
        (error != null ? <ErrorComponent msg={error} /> : page);
}
export default HomePage;